# InsightsPulse
Retrieve sports data from third party apis process, organize, analyze them and store insights.

## Setup Locally
1. Copy `InsightsPulse\examples\.env.example` to `InsightsPulse\`
   1. Rename it to **.env**
2. Copy `InsightsPulse\examples\config.example.yml` to `InsightsPulse\`
   1. Rename it to **config.yml**
3. Fill Both config files based on the local setup
4. **RUN** `> docker compose build`
5. **RUN** `> docker compose up`
   
## Deploy

> Deploy it to Amazon **Elastic Container Registry**(ECS)

1. Use production config variables(check that is correct)
2. Authenticate using **aws command line**
3. Create docker image 
    `> docker build -t imagename .`
4. Add Tag
   `> docker tag imagename:latest {{amazon_endpoint}}/{{repository_name}}:latest`
5. Push Image to Aws
   `> docker tag imagename:latest {{amazon_endpoint}}/{{repository_name}}:latest`
