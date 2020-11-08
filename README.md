# weatherservice


## Notes 
*Please do not supply your name or email address in this document. We're doing our best to remain unbiased.* 
some things I would implement if I had more time:

- app getting deployed onto heroku took a bit longer than I would have liked but mostly due to having not used it in awhile. Wasn't bad once I got through the finnicky gotchas (like them assigning their own port binding... which makes some sense)
- chose to build simple webserver from go and from one main.go file because of the simplicity in the server. If there was more functionality, I likely would have organized the repo a bit differently with more folders
)
- chose to use vue because vue-cli sets me up quickly for the few options of login/registration and home page; not many reuseable components with data passing shown but I tried at least passing some props into the header image for the asset image. also using a FE component framework allows for the project to grow easier if/when there are more things into play 

### Date 
Saturday Nov 7, 2020 


### Location of deployed application 
https://whattheweather.herokuapp.com/ (it is a free dyno so may take some time to load) 
### Time spent 
~5 
How much time did you spend on the assignment? Normally, this is expressed in hours. 
 
### Assumptions made 
Use this section to tell us about any assumptions that you made when creating your solution. 

- minicache implementation: assuming I'm not going to be supported millions of requests of different cities within 5 min (otherwise memory would balloon)
- cities entered is exactly what the user wanted. I started putting in a google maps api in place but figured it may take more time to configure properly (started process of using some vue libraries with specified place but parsing out the place details returned for the city name was a pain) 
- that users won't care about needing to delete their account once created

### Shortcuts/Compromises made 
If applicable. Did you do something that you feel could have been done better in a real-world application? Please 
let us know. 
 
- fixing my docker build file to build the FE section and place it into the new container (for some reason it kept running out of memory so I ended up just building locally and then placing the built files into the container


### Stretch goals attempted 
If applicable, use this area to tell us what stretch goals you attempted. What went well? What do you wish you 
could have done better? If you didn't attempt any of the stretch goals, feel free to let us know why. 

- simple UI
- authentication using firebase
- deploying app to heroku
- proxy weather api + quick mini cache implementation
 
### Instructions to run assignment locally 
If applicable, please provide us with the necessary instructions to run your solution. 
it's deployed on heroku but to run locally:

cd server
go build
./src

cd ../client
npm run serve
 
I believe it will run at 8080 for the server and npm on whichever next port is avail. the client shoudl be proxied to call server at 8080


### What did you not include in your solution that you want us to know about? 
Were you short on time and not able to include something that you want us to know 
about? Please list it here so that we know that you considered it. 
 
- the google places auto complete search; I had a working version but honestly FE takes me longer and I wanted to work within the specified time scope while allowing time for configuring deployment etc. This is I feel a much better solution and would get rid of the assumption that whatever city is typed in is the one the user wants (ie. Vancouver BC instead of Vancouver WA). I spent about 30 min getting a semi working version getting the lat long and having the google api auto complete/predict the address but then worried I wouldn't finish the initial task of entering in a city. 

### Other information about your submission that you feel it's important that we know if applicable. 
I enjoy backend/devopsy things more than FE although can be open to both since I understand this is also a fullstack position...Looking back at today, I semi wish i spent more time implementing different features in the server but found myself setting up the client much longer.   

### Your feedback on this technical challenge 
Decent challenge which got me back to quick hacking for a bit and sent me back to heroku. Would have liked challenges that had a bit more complexity vs tedious/setup but that was probably coupd have been up to me to decide on my own :)