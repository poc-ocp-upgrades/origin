package printers

import (
 "io"
 "text/tabwriter"
)

const (
 tabwriterMinWidth = 6
 tabwriterWidth    = 4
 tabwriterPadding  = 3
 tabwriterPadChar  = ' '
 tabwriterFlags    = 0
)

func GetNewTabWriter(output io.Writer) *tabwriter.Writer {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return tabwriter.NewWriter(output, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
}
