# my-project
1. Running `docker-compose up` will start the app, and you can 
access the app using localhost:8080/ which will give you 
grpahql playground and you can give the query in the below format
```
query ingestData {
  ingestData {
    ingestionStatus
  }
}
```

2. If you want to use your api-key then change the Dockerfile to use your own api-key and do the step 1.

3. To run in k8s, install minikube and run the below command which will create the pods, services.
```
kubectl apply -f graphql-networkpolicy.yaml,database-claim0-persistentvolumeclaim.yaml,database-deployment.yaml,database-service.yaml,server-deployment.yaml,server-service.yaml
```

4. To access the app from localhost you need to port-forward the server service, use the below command
```
kubectl port-forward service/server 8080:8080
```

5. To check if the data has persisted in the postgres, port-forward the postgres port too, use the below command
```
kubectl port-forward service/database 5432:5432
```

6. Use the same credentials in pgadmin client that are avaiable in the database-deployment.yaml or you can use your own, please make sure you use the same env variables in both deployment files.
