This repository serves as a POC for operator dependency.

## Motive
The idea is to declare a dependency of one operator on a CRD provided by another
operator and installation of first operator should install the second operator.

The dependency is declared using the `required CRD` block. The depedency in this
POC is declared in
https://github.com/jarifibrahim/two-operators/blob/master/app-operator/deploy/olm-catalog/app-operator/0.1.0/app-operator.v0.1.0.clusterserviceversion.yaml#L18

## How to enable operators
1. Login as admin
```
oc login -u system:admin
```
2. Install OLM using
```
oc create -f https://raw.githubusercontent.com/operator-framework/operator-lifecycle-manager/master/deploy/upstream/quickstart/olm.yaml
```

Please ensure all the pods in the OLM namepsace are running before running the
following commands.
You can watch the status of all pods in `olm` namespace using
```
oc get pods -n olm -w
```

3. Enable both the operators by creating a new catalogsource.
```
oc create -f catalogSource/configmap_devopsconsole.yaml
```
This will create a new configmap in the olm namespace.
Now run
```
oc create -f catalogSource/devops-console-catalog.yaml
```
This will create a new catalogSource and the `app-operator` and
`dependant-operator` should be visible in the operator management page (You'll need
openshift 4 console to see this.
Run https://github.com/operator-framework/operator-lifecycle-manager/blob/master/scripts/run_console_local.sh
to run openshift 4 console locally)

Verify you have a catalogsource using 
```
oc get catalogsources -n olm
```
This should have `devops-console-catalog` source.

## How to install the app-operator
Create a subscription in the `myproject` namespace. Change the namespace in the
`subscription.yaml` to use a different namespace
```
oc create -f ./subscription.yaml
```

This will create a subscription and an install plan. The install plan will
contain the operator dependency and OLM will resolve the dependency by
installing the dependant operator along with the app-operator.

Verify the subscription is created using
```
oc get subscriptions.operators.coreos.com --all-namespaces
```

Verify an install plan is created
```
oc get installplans.operators.coreos.com --all-namespaces
```
