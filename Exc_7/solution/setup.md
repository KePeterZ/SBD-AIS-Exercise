Since we could not get swarm done in a multi-computer cluster, I chose to do a single worker implementation (as written in the email).

To init the cluster, I did "docker swarm init", and to deploy, I did "docker stack deploy -c docker-compose.yml sbd3ue". However, I would use Kubernetes even with a gun to my head.

also this does not work on my machine since I'm using an Apple M1 Pro architecture, tested on an x86 machine though, it worked