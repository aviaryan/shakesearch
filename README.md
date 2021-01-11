# ShakeSearch

Welcome to the Pulley Shakesearch Take-home Challenge! In this repository,
you'll find a simple web app that allows a user to search for a text string in
the complete works of Shakespeare.

You can see a live version of the app at
https://pulley-shakesearch.herokuapp.com/. Try searching for "Hamlet" to display
a set of results.

In it's current state, however, the app is just a rough prototype. The search is
case sensitive, the results are difficult to read, and the search is limited to
exact matches.

## Your Mission

Improve the search backend. Think about the problem from the **user's perspective**
and prioritize your changes according to what you think is most useful. 

## Evaluation

We will be primarily evaluating based on how well the search works for users. A search result with a lot of features (i.e. multi-words and mis-spellings handled), but with results that are hard to read would not be a strong submission. 


## Submission

1. Fork this repository and send us a link to your fork after pushing your changes. 
2. Heroku hosting - The project includes a Heroku Procfile and, in its
current state, can be deployed easily on Heroku's free tier.
3. In your submission, share with us what changes you made and how you would prioritize changes if you had more time.


## Upgrades

- [x] Add option for case-insensitive search ("hamlet").
- [x] Highlight search result in web app to end-user.
- [x] Return book/play name with the search result.
- [ ] Add option for word search.
- [x] Show search results cleanly in table with book/play name, search results count, etc.
- [x] Trim context information around the searched phrase if it goes out of context of the book/play (try searching `water cools`). 
- [ ] Add loader on UI to show it's searching.
- [x] Prevent searches like "GUTENBERG" from showing in results.
- [x] Fix bug where string slicing goes out of bound and an application error happens (at the end).
- [ ] Unicode breaking on the page, fix it.
