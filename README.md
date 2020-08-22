# Signavio

The 1st application is server which returns ```{"id": "1", "message": "Hello world"}``` upon a GET request by a HTTP client and has host name
```signavio-server.com```.

The 2nd application is a client which makes a request to server and reverse the message fetched in response as follows
```world Hello```

Both applications have Dockerfile in their respective directories. Docker images are built and have been pushed to hub.docker.com
For server: ```docker pull hari95kishore/signavio```
For client: ```docker pull hari95kishore/signavio-client```

I have used these images in respective helm values.yaml.

I have creaed helm charts for both applications and can be deployed as follows
For deploying server into minikube:
```
minikube start
cd server
helm upgrade --install signavio signavioserver/ --values signavioserver/values.yaml
```

Helm upgrade --install will look for existing release signavio and upgrade it, if there is no such releasename, it creates a new one.

Server has 2 replicas and is set up as ```service.Type: NodePort``` so that it can be accessed outside cluster.

Nginx ingress controller is used and host ```signavio-server.com``` is set in ingress section of values.yaml.

Once successfully deployed, you can see that hitting ```http://signavio-server.com``` in browser doesn't return json as your local machine can't resolve the name.

Run ```minikube ip``` and add ```<minikube_ip_address> signavio-server.com``` in your /etc/hosts file so that your machine can resolve can resolve it.

Now hit ```http://signavio-server.com``` in browser and you should see response json ```{"id": "1", "message": "Hello world"}```

Now that server is up and running, lets move on to client deployment.
For deploying client in minikube:
```
cd client
helm upgrade --install signavioclient signavioclient/ --values signavioclient/values.yaml
```

Once the deployment is successful and container completed running, you can look for the result as follows:
```
kubectl logs -f <podname>
```

and you should be able to see the result ```world Hello``` reverse message of server response.

Please note that client service is deployed as ```Kind: Deployment``` with ```restartPolicy: OnFailure```. Unfortunately as of now Kubernetes seems to honour restartPolicy only for ```Kind: Job``` and batch but doesn't honour for deployment kind.

As a result restartPolicy: Always gets applied and pods keep restarting.

Here is the open issue https://github.com/kubernetes/kubernetes/issues/24725 and enhancement https://github.com/kubernetes/enhancements/pull/912 being tracked.

Hope everything runs smoothly!

P.S. : I really had fun and enjoyed the challenge. Looking forward for the feedback!


