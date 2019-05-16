package node

import (
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	bootstraputil "k8s.io/cluster-bootstrap/token/util"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
)

func CreateNewTokens(client clientset.Interface, tokens []kubeadmapi.BootstrapToken) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return UpdateOrCreateTokens(client, true, tokens)
}
func UpdateOrCreateTokens(client clientset.Interface, failIfExists bool, tokens []kubeadmapi.BootstrapToken) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, token := range tokens {
		secretName := bootstraputil.BootstrapTokenSecretName(token.Token.ID)
		secret, err := client.CoreV1().Secrets(metav1.NamespaceSystem).Get(secretName, metav1.GetOptions{})
		if secret != nil && err == nil && failIfExists {
			return errors.Errorf("a token with id %q already exists", token.Token.ID)
		}
		updatedOrNewSecret := token.ToSecret()
		err = apiclient.TryRunCommand(func() error {
			if err := apiclient.CreateOrUpdateSecret(client, updatedOrNewSecret); err != nil {
				return errors.Wrapf(err, "failed to create or update bootstrap token with name %s", secretName)
			}
			return nil
		}, 5)
		if err != nil {
			return err
		}
	}
	return nil
}
