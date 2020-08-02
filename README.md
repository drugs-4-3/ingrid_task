# ingrid_task

## To launch: 
1. `git clone git clone git@github.com:drugs-4-3/ingrid_task.git`
2. `cd ingrid_task/`
3. `docker build . -t ingrid_test_build`
4. `docker run -p 8080:8080 --env-file .env ingrid_test_build`

The service will run on port 8080. If you want to use different port - change it in .env and Dockerfile 
