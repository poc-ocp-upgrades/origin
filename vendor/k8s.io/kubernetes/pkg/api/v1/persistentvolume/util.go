package persistentvolume

import (
	goformat "fmt"
	corev1 "k8s.io/api/core/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func getClaimRefNamespace(pv *corev1.PersistentVolume) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pv.Spec.ClaimRef != nil {
		return pv.Spec.ClaimRef.Namespace
	}
	return ""
}

type Visitor func(namespace, name string, kubeletVisible bool) (shouldContinue bool)

func VisitPVSecretNames(pv *corev1.PersistentVolume, visitor Visitor) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	source := &pv.Spec.PersistentVolumeSource
	switch {
	case source.AzureFile != nil:
		if source.AzureFile.SecretNamespace != nil && len(*source.AzureFile.SecretNamespace) > 0 {
			if len(source.AzureFile.SecretName) > 0 && !visitor(*source.AzureFile.SecretNamespace, source.AzureFile.SecretName, true) {
				return false
			}
		} else {
			if len(source.AzureFile.SecretName) > 0 && !visitor(getClaimRefNamespace(pv), source.AzureFile.SecretName, true) {
				return false
			}
		}
		return true
	case source.CephFS != nil:
		if source.CephFS.SecretRef != nil {
			ns := getClaimRefNamespace(pv)
			if len(source.CephFS.SecretRef.Namespace) > 0 {
				ns = source.CephFS.SecretRef.Namespace
			}
			if !visitor(ns, source.CephFS.SecretRef.Name, true) {
				return false
			}
		}
	case source.Cinder != nil:
		if source.Cinder.SecretRef != nil && !visitor(source.Cinder.SecretRef.Namespace, source.Cinder.SecretRef.Name, true) {
			return false
		}
	case source.FlexVolume != nil:
		if source.FlexVolume.SecretRef != nil {
			ns := getClaimRefNamespace(pv)
			if len(source.FlexVolume.SecretRef.Namespace) > 0 {
				ns = source.FlexVolume.SecretRef.Namespace
			}
			if !visitor(ns, source.FlexVolume.SecretRef.Name, true) {
				return false
			}
		}
	case source.RBD != nil:
		if source.RBD.SecretRef != nil {
			ns := getClaimRefNamespace(pv)
			if len(source.RBD.SecretRef.Namespace) > 0 {
				ns = source.RBD.SecretRef.Namespace
			}
			if !visitor(ns, source.RBD.SecretRef.Name, true) {
				return false
			}
		}
	case source.ScaleIO != nil:
		if source.ScaleIO.SecretRef != nil {
			ns := getClaimRefNamespace(pv)
			if source.ScaleIO.SecretRef != nil && len(source.ScaleIO.SecretRef.Namespace) > 0 {
				ns = source.ScaleIO.SecretRef.Namespace
			}
			if !visitor(ns, source.ScaleIO.SecretRef.Name, true) {
				return false
			}
		}
	case source.ISCSI != nil:
		if source.ISCSI.SecretRef != nil {
			ns := getClaimRefNamespace(pv)
			if len(source.ISCSI.SecretRef.Namespace) > 0 {
				ns = source.ISCSI.SecretRef.Namespace
			}
			if !visitor(ns, source.ISCSI.SecretRef.Name, true) {
				return false
			}
		}
	case source.StorageOS != nil:
		if source.StorageOS.SecretRef != nil && !visitor(source.StorageOS.SecretRef.Namespace, source.StorageOS.SecretRef.Name, true) {
			return false
		}
	case source.CSI != nil:
		if source.CSI.ControllerPublishSecretRef != nil {
			if !visitor(source.CSI.ControllerPublishSecretRef.Namespace, source.CSI.ControllerPublishSecretRef.Name, false) {
				return false
			}
		}
		if source.CSI.NodePublishSecretRef != nil {
			if !visitor(source.CSI.NodePublishSecretRef.Namespace, source.CSI.NodePublishSecretRef.Name, true) {
				return false
			}
		}
		if source.CSI.NodeStageSecretRef != nil {
			if !visitor(source.CSI.NodeStageSecretRef.Namespace, source.CSI.NodeStageSecretRef.Name, true) {
				return false
			}
		}
	}
	return true
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
