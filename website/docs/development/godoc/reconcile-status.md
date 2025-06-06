<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# progress

```go
import "github.com/k-orc/openstack-resource-controller/v2/internal/controllers/generic/progress"
```

## Index

- [type ReconcileStatus](<#ReconcileStatus>)
  - [func NeedsRefresh\(\) ReconcileStatus](<#NeedsRefresh>)
  - [func NewReconcileStatus\(\) ReconcileStatus](<#NewReconcileStatus>)
  - [func WaitingOnFinalizer\(finalizer string\) ReconcileStatus](<#WaitingOnFinalizer>)
  - [func WaitingOnObject\(kind, name string, waitingOn WaitingOnEvent\) ReconcileStatus](<#WaitingOnObject>)
  - [func WaitingOnOpenStack\(waitingOn WaitingOnEvent, pollingPeriod time.Duration\) ReconcileStatus](<#WaitingOnOpenStack>)
  - [func WrapError\(err error\) ReconcileStatus](<#WrapError>)
  - [func \(r ReconcileStatus\) GetError\(\) error](<#ReconcileStatus.GetError>)
  - [func \(r ReconcileStatus\) GetProgressMessages\(\) \[\]string](<#ReconcileStatus.GetProgressMessages>)
  - [func \(r ReconcileStatus\) GetRequeue\(\) time.Duration](<#ReconcileStatus.GetRequeue>)
  - [func \(r ReconcileStatus\) NeedsRefresh\(\) ReconcileStatus](<#ReconcileStatus.NeedsRefresh>)
  - [func \(r ReconcileStatus\) NeedsReschedule\(\) \(bool, error\)](<#ReconcileStatus.NeedsReschedule>)
  - [func \(r ReconcileStatus\) Return\(log logr.Logger\) \(ctrl.Result, error\)](<#ReconcileStatus.Return>)
  - [func \(r ReconcileStatus\) WaitingOnFinalizer\(finalizer string\) ReconcileStatus](<#ReconcileStatus.WaitingOnFinalizer>)
  - [func \(r ReconcileStatus\) WaitingOnObject\(kind, name string, waitingOn WaitingOnEvent\) ReconcileStatus](<#ReconcileStatus.WaitingOnObject>)
  - [func \(r ReconcileStatus\) WaitingOnOpenStack\(waitingOn WaitingOnEvent, pollingPeriod time.Duration\) ReconcileStatus](<#ReconcileStatus.WaitingOnOpenStack>)
  - [func \(r ReconcileStatus\) WithError\(err error\) ReconcileStatus](<#ReconcileStatus.WithError>)
  - [func \(r ReconcileStatus\) WithProgressMessage\(msgs ...string\) ReconcileStatus](<#ReconcileStatus.WithProgressMessage>)
  - [func \(r ReconcileStatus\) WithReconcileStatus\(o ReconcileStatus\) ReconcileStatus](<#ReconcileStatus.WithReconcileStatus>)
  - [func \(r ReconcileStatus\) WithRequeue\(requeue time.Duration\) ReconcileStatus](<#ReconcileStatus.WithRequeue>)
- [type WaitingOnEvent](<#WaitingOnEvent>)


<a name="ReconcileStatus"></a>
## type [ReconcileStatus](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L38>)

ReconcileStatus represents the status of the current reconcile.

nil is a valid ReconcileStatus. It is safe to call all methods on a nil ReconcileStatus.

You MUST use the return value of any method which returns ReconcileStatus: the returned object may not be equal to the receiver, for example when the receiver is nil.

```go
type ReconcileStatus = *reconcileStatus
```

<a name="NeedsRefresh"></a>
### func [NeedsRefresh](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L263>)

```go
func NeedsRefresh() ReconcileStatus
```

NeedsRefresh is a convenience method which returns a new ReconcileStatus with NeedsRefresh.

<a name="NewReconcileStatus"></a>
### func [NewReconcileStatus](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L48>)

```go
func NewReconcileStatus() ReconcileStatus
```

NewReconcileStatus returns an empty ReconcileStatus

<a name="WaitingOnFinalizer"></a>
### func [WaitingOnFinalizer](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L227>)

```go
func WaitingOnFinalizer(finalizer string) ReconcileStatus
```

WaitingOnFinalizer is a convenience method which returns a new ReconcileStatus with WaitingOnFinalizer.

<a name="WaitingOnObject"></a>
### func [WaitingOnObject](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L217>)

```go
func WaitingOnObject(kind, name string, waitingOn WaitingOnEvent) ReconcileStatus
```

WaitingOnObject is a convenience method which returns a new ReconcileStatus with WaitingOnObject.

<a name="WaitingOnOpenStack"></a>
### func [WaitingOnOpenStack](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L251>)

```go
func WaitingOnOpenStack(waitingOn WaitingOnEvent, pollingPeriod time.Duration) ReconcileStatus
```

WaitingOnOpenStack is a convenience method which returns a new ReconcileStatus with WaitingOnOpenStack.

<a name="WrapError"></a>
### func [WrapError](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L53>)

```go
func WrapError(err error) ReconcileStatus
```

WrapError returns a ReconcileStatus containing the given error

<a name="ReconcileStatus.GetError"></a>
### func \(ReconcileStatus\) [GetError](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L69>)

```go
func (r ReconcileStatus) GetError() error
```

GetError returns an error representing all errors which have been added to this ReconcileStatus. If multiple errors have been added they will have been combined with errors.Join\(\)

<a name="ReconcileStatus.GetProgressMessages"></a>
### func \(ReconcileStatus\) [GetProgressMessages](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L58>)

```go
func (r ReconcileStatus) GetProgressMessages() []string
```

GetProgressMessages returns all progress messages which have been added to this ReconcileStatus

<a name="ReconcileStatus.GetRequeue"></a>
### func \(ReconcileStatus\) [GetRequeue](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L79>)

```go
func (r ReconcileStatus) GetRequeue() time.Duration
```

GetRequeue returns the time after which the current object should be reconciled again. A value of 0 indicates that no requeue is requested.

<a name="ReconcileStatus.NeedsRefresh"></a>
### func \(ReconcileStatus\) [NeedsRefresh](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L258>)

```go
func (r ReconcileStatus) NeedsRefresh() ReconcileStatus
```

NeedsRefresh indicates that the resource status needs to be refreshed. It sets an appropriate progress message and ensures that the object will be reconciled again immediately.

<a name="ReconcileStatus.NeedsReschedule"></a>
### func \(ReconcileStatus\) [NeedsReschedule](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L94>)

```go
func (r ReconcileStatus) NeedsReschedule() (bool, error)
```

NeedsReschedule returns a boolean value indicating whether the ReconcileStatus will set the Progressing condition to true, and therefore that we intend to be scheduled again. It additionally returns any error associated with the ReconcileStatus.

NeedsReschedule is used to shortcut reconciliation if any precondition has not been met.

<a name="ReconcileStatus.Return"></a>
### func \(ReconcileStatus\) [Return](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L108>)

```go
func (r ReconcileStatus) Return(log logr.Logger) (ctrl.Result, error)
```

Return returns the the \(ctrl.Result, error\) expected by controller\-runtime for a ReconcileStatus.

If a ReconcileStatus contains a TerminalError, Return will log the error directly instead of returning it to controller\-runtime, as this would cause an undesirable reschedule.

<a name="ReconcileStatus.WaitingOnFinalizer"></a>
### func \(ReconcileStatus\) [WaitingOnFinalizer](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L222>)

```go
func (r ReconcileStatus) WaitingOnFinalizer(finalizer string) ReconcileStatus
```

WaitingOnFinalizer adds a progress message indicating that we are waiting for a specific finalizer to be removed.

<a name="ReconcileStatus.WaitingOnObject"></a>
### func \(ReconcileStatus\) [WaitingOnObject](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L203>)

```go
func (r ReconcileStatus) WaitingOnObject(kind, name string, waitingOn WaitingOnEvent) ReconcileStatus
```

WaitingOnObject adds a progress message indicating that we are waiting on a kubernetes object of type kind with name. We expect the controller to have an appropriate watch and handler for this event, so WaitingOnObject does not add an explicit requeue.

<a name="ReconcileStatus.WaitingOnOpenStack"></a>
### func \(ReconcileStatus\) [WaitingOnOpenStack](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L235>)

```go
func (r ReconcileStatus) WaitingOnOpenStack(waitingOn WaitingOnEvent, pollingPeriod time.Duration) ReconcileStatus
```

WaitingOnOpenStack indicates that we are waiting for an event on the current OpenStack resource. It adds an appropriate progress message. It also adds a requeue with the requested polling period, as we are not able to receive triggers for OpenStack events.

<a name="ReconcileStatus.WithError"></a>
### func \(ReconcileStatus\) [WithError](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L161>)

```go
func (r ReconcileStatus) WithError(err error) ReconcileStatus
```

WithError returns a ReconcileStatus containing the given error joined to any existing errors.

<a name="ReconcileStatus.WithProgressMessage"></a>
### func \(ReconcileStatus\) [WithProgressMessage](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L128>)

```go
func (r ReconcileStatus) WithProgressMessage(msgs ...string) ReconcileStatus
```

WithProgressMessage returns a ReconcileStatus with the given progress messages in addition to any already present.

<a name="ReconcileStatus.WithReconcileStatus"></a>
### func \(ReconcileStatus\) [WithReconcileStatus](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L175>)

```go
func (r ReconcileStatus) WithReconcileStatus(o ReconcileStatus) ReconcileStatus
```

WithReconcileStatus returns a ReconcileStatus combining all properties of the given ReconcileStatus.

<a name="ReconcileStatus.WithRequeue"></a>
### func \(ReconcileStatus\) [WithRequeue](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L144>)

```go
func (r ReconcileStatus) WithRequeue(requeue time.Duration) ReconcileStatus
```

WithRequeue returns a ReconcileStatus with a request to requeue after the given time. If the ReconcileStatus already requests a requeue, the returned object will have the lesser of the existing and requested requeues.

<a name="WaitingOnEvent"></a>
## type [WaitingOnEvent](<https://github.com/k-orc/openstack-resource-controller/blob/main/internal/controllers/generic/progress/reconcile_status.go#L186>)

WaitingOnEvent represents the type of event we are waiting on

```go
type WaitingOnEvent int
```

<a name="WaitingOnCreation"></a>

```go
const (
    // WaitingOnCreation indicates waiting for an object to be created
    WaitingOnCreation WaitingOnEvent = iota

    // WaitingOnReady indicates that an object exists but is not yet in the necessary state
    WaitingOnReady

    // WaitingOnDeletion indicates waiting for an object to be deleted
    WaitingOnDeletion
)
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
