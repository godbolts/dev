# Match-Me Web Application

## Setup

Match-Me uses a PostgeSQL database, if you do not have the PostgeSQL installed on your machine then you can download it from https://www.postgresql.org/download/. I have Set up my server to listen on port **5432**. When choosing to add packages then you need to add PostGIS Sparcial Extention as well, it is nessesary for GPS data.

In your PostgreSQL setup you need to set up a superuser with a password, you need to choose a password and whatever password you choose, you must update the SUPER_USER_PASS variable in the **config.env** file in the server folder - this allows the server to create a user and a database.

I have added a multitude of files to make the application easy to start on the three operation systems. The **run_project.js** file runs all the appropriate files by identifing the operation system and calling the respective command files, the files ending in *.sh* and *.bat* files depending on the system. Just in case I also added system specific parallels. Bat files run on Windows systems and feed into the command prompt, sh files run in Linux/Mac systems and feed into their command lines.



## Database Setup

npm install pg