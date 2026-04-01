package v1

import "github.com/robinlg/iamlib/pkg/scheme"

// GroupName is the group name use in this package.
// If use a public domain name, need set the GroupName to service name.
// For example: if restful path is: https://robinlg.com/apimachinery/v1/secrets, we can set GroupName="apimachinery".
const GroupName = "iam.authz"

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = scheme.GroupVersion{Group: GroupName, Version: "v1"}
