# crdb_connection_testing

## Requirements:
- Setup roachprod [roachprod](https://github.com/cockroachdb/cockroach/tree/master/pkg/cmd/roachprod)
  - Note: roachprod is only available to Cockroach Labs employees
- Optional: setup GCE worker (also only applicable to Cockroach Labs employees)
- Install go, this tutorial assumes you have go setup.
- Install [pgx](https://github.com/jackc/pgx) and make sure it is in your $GOPATH

1. Create roachprod cluster
  - Use script to setup roachprod cluster faster
    - Example: /start_prod_cluster.sh 3 $CLUSTER_NAME v21.1.0-beta.1
    - Note: assumes roachprod is in your $PATH
2. Run Go script to spin up connections.
  - go run connection_script.go \`{roachprod pgurl --external $CLUSTER_NAME:1-2}\`
  - Leave one node alone to use as baseline reference. Do not pass URL to script.
  - The previous command assumes you have roachprod installed on your machine
  - otherwise, grab the urls using to the roachprod machines using `roachprod pgurl --external $CLUSTER_NAME`
  - Tip: if you start running into socket errors, check your ulimit -n on the GCE worker / locally
3. Open Admin UI
  - roachprod adminurl $CLUSTER_NAME
4. Under Metrics -> SQL verify that the connections have been created
![image](https://user-images.githubusercontent.com/25163644/112694581-00b49400-8e59-11eb-9b3a-7e20b56f3300.png)
5. Under Advanced Debug -> Heap and Advanced Debug -> Profile: check the CPU / Memory usage.
![image](https://user-images.githubusercontent.com/25163644/112694670-1de96280-8e59-11eb-9d4c-00984ebd54ab.png)
