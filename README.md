# Connect-DB-Demo [MongoDB]

I prefer to use IBM [CLI](https://github.com/IBM-Cloud/ibm-cloud-cli-release/releases/) to demo, of course, you could get all information that I will mention from instance dashboard.

### Provision one instance 
> skip this step if you already have one

1. To create a database, you’ll need to log in first. If you’re using a federated identity, you’ll use:
```
ibmcloud login --sso
```
2. Create your databases
```
ibmcloud resource service-instance-create <instance_name> databases-for-mongodb standard us-south
```
for example:
```
ibmcloud resource service-instance-create Demo-MongoDB databases-for-mongodb standard us-south
```
### Get connection string
1. Once the database is done provisioning, you can get the credentials using the cloud-databases plugin. Install that using the following command
```
ibmcloud plugin install cloud-databases
```
2. Get ```hostname```, ```port```, ```database```, ```auth source``` via CLI and set them as environment variables
```
ibmcloud cdb cxn Demo-MongoDB
Retrieving public connection strings for Demo-MongoDB...
OK

Type      Connection String
MongoDB   mongodb://admin:$PASSWORD@ba257b06-aa41-4b69-9e92-7e76cd2f578c-0.bkr06mid0v493nkn6i3g.databases.appdomain.cloud:31742,ba257b06-aa41-4b69-9e92-7e76cd2f578c-1.bkr06mid0v493nkn6i3g.databases.appdomain.cloud:31742/ibmclouddb?authSource=admin&replicaSet=replset
CLI       mongo -u admin -p $PASSWORD --ssl --sslCAFile 359541d4-b7a6-11e9-950e-fefc37a38a5a --authenticationDatabase admin --host replset/ba257b06-aa41-4b69-9e92-7e76cd2f578c-0.bkr06mid0v493nkn6i3g.databases.appdomain.cloud:31742,ba257b06-aa41-4b69-9e92-7e76cd2f578c-1.bkr06mid0v493nkn6i3g.databases.appdomain.cloud:31742
```
```
export HOSTNAME=ba257b06-aa41-4b69-9e92-7e76cd2f578c-0.bkr06mid0v493nkn6i3g.databases.appdomain.cloud
export PORT=31742
export DATABASE=ibmclouddb
export AUTHSOURCE=admin
```
3. Set password for database
```
ibmcloud cdb deployment-user-password Demo-MongoDB  admin <your_password>
The user's password is being changed with this task:

Key                   Value
ID                    crn:v1:bluemix:public:databases-for-mongodb-preproduction:us-south:a/b2742234a9bf412a8183644a5a92cd95:ba257b06-aa41-4b69-9e92-7e76cd2f578c:task:acd40447-138a-492d-a034-88980dda0bfd
Deployment ID         crn:v1:bluemix:public:databases-for-mongodb-preproduction:us-south:a/b2742234a9bf412a8183644a5a92cd95:ba257b06-aa41-4b69-9e92-7e76cd2f578c::
Description           Updating user.
Created At            2019-09-05T09:32:48Z
Status                running
Progress Percentage   0

Status                completed
Progress Percentage   100
Location              https://api.preproduction.us-south.databases.cloud.ibm.com/v4/ibm/deployments/crn:v1:bluemix:public:databases-for-mongodb-preproduction:us-south:a%2Fb2742234a9bf412a8183644a5a92cd95:ba257b06-aa41-4b69-9e92-7e76cd2f578c::
OK
```
```
export USERNAME=admin
export PASSWORD=<your_password>
```
4. You’ll also need to decode the CA certificate that your databases need for authentication. To decode it, run the following command then make sure to copy the decoded certificate and save it to a file on your system
```
ibmcloud cdb cacert Demo-MongoDB
```
```
export CAFILE=<your_cafile_path>
```
### Test
1. Run it, and it will show the inserted ID for one new record
```
make run
```
```
ObjectID("5d70e3b1ac612c5a21f9e7dd")
```
