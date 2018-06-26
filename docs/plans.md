Current goal
I want to be able to have a front end that shows a teams ranking compared to all the other teams in their league.

A team belongs to a league and has a rank that is unique within that league.

Next possible steps:
- Use persistent storage such as sqllite or something running in a docker image such as a NoSQL database
- Acually mimic a real data structure of a person's team in a League
- Create 404 page
- Anything from this book https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/08.1.html

Goals:
- Enjoy programming
- Learn new skills
- Have a finished product
- Don't give up

Accomplishments:
- Created makefile
- Set up new computer
- Wrote tests
- Create static front-end with styling, etc.
- Make MVC
- Refactor load template
- Create logger using go routine
- Serve static files like Favicon


Data model
A user has many user teams
    - Only one can be active at a time
A user is part of a user league
    - A user league has many user teams