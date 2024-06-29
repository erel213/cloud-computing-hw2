# cloud-computing-hw2

## Data Modeling

![data-modeling](https://github.com/erel213/cloud-computing-hw2/assets/99283653/18a7051f-1a2b-4866-add2-3cd8029ee2a6)

## Deploy Instructions 
1. To deploy the cloud resources, run the init.sh script.
2. Connect to the bastion machine using the generated __bastion_host_key__ from init.sh.
3. Clone the code repository on the bastion machine.
4. Create the database using psql from the bastion machine.
5. Perform the following commands on the bastion machine to use go migrate for database schema migration [text](https://github.com/golang-migrate/migrate):
    a. Download migrate.linux-amd64.tar.gz from https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz
    b. Move the migrate binary to /usr/local/bin/ by running sudo mv ~/go/bin/migrate /usr/local/bin/
    c. Verify the installation of migrate by running migrate -version.
6. Execute the following command to perform the migrations:
    migrate -path cloud-computing-hw2/internal/infrastructure/db/migrations -database "postgres://<username>:<password>@<host>:<port>/<dbname>?sslmode=disable" up
#### READY TO GO !

## Architecture Disscussion
The application is deployed on AWS using a 3-tier architecture. The presentation layer consists of a load balancer and a bastion host, which is used to set up the database. The bastion host can communicate with both the application layer and the data layer through a NAT gateway.

The application layer is comprised of an ECS cluster that utilizes Fargate for deployment. Ingress traffic to the application layer is only allowed from the presentation layer.

The data layer consists of an RDS instance running on the PostgreSQL engine. Ingress traffic to the data layer is permitted from both the application layer and the bastion host.

## Scaling and Cost Analysis

### Scaling
The system is designed to handle increased user load by utilizing AWS services such as ECS and Fargate. ECS allows for easy scaling of the application layer by automatically provisioning and managing containers. Fargate, on the other hand, abstracts away the underlying infrastructure, allowing for seamless scaling based on demand.

As the number of users and load increases, ECS can automatically scale up the number of Fargate tasks to handle the additional traffic. This ensures that the application layer can handle the increased load without any manual intervention. Additionally, the load balancer distributes the incoming requests evenly across the available tasks, further optimizing the system's performance.

In terms of the data layer, the design choices can also impact scaling. For example, if the system relies on a single RDS instance, it may become a bottleneck as the user load increases. In such cases, sharding the database or utilizing read replicas can help distribute the load and improve scalability.

### Cost Analysis
When it comes to cost, scaling can have an impact on the overall expenses. As the number of Fargate tasks increases, so does the cost of running those tasks. However, AWS provides cost optimization features such as auto-scaling policies and spot instances that can help mitigate the expenses.

By setting up auto-scaling policies, you can define rules to automatically scale the number of Fargate tasks based on metrics like CPU utilization or request count. This ensures that the system scales up only when necessary, minimizing unnecessary costs during periods of low traffic.

Additionally, spot instances can be utilized to further reduce costs. Spot instances are spare EC2 instances that are available at a significantly lower price compared to on-demand instances. By leveraging spot instances for the ECS cluster, you can achieve cost savings while maintaining the required level of scalability.

In terms of the data layer, the choice of database engine and instance type can also impact costs. For example, using a managed database service like Amazon Aurora can provide cost savings compared to running a self-managed database on EC2 instances.

Overall, the system's reaction to increased users and load is efficient scaling through ECS and Fargate, while cost optimization can be achieved through auto-scaling policies, spot instances, and smart choices in the data layer design.
