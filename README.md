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

- [x] Add <s>option</s>feature for case-insensitive search ("hamlet").
- [x] Highlight search query occurances on the web app (server-side highlighting via `<mark>` works in this case).
- [x] Return book/play name with the search result.
- [x] Show search results in a table with book/play name, search results count, etc.
- [x] Add `<br>` in search results to make them readable and helpful in getting context.
- [x] Add scroll-to-top button to return to the search bar.
- [x] Make it usable on mobile.
- [x] Trim context information around the searched phrase if it goes out of context of the book/play (try searching `water cools`). 
- [x] Add a loading message in UI when searching.
- [x] Return results in order of the story so that end-user can follow filter results and make sense of the story.
- [x] Prevent searches like "GUTENBERG" from showing in results.
- [x] Fix bug where string slicing goes out of bound and an application error happens (at the end).
- [x] Fix very-similar results which happens if search query is present many times near one another (try searching `hamlet`, then Cmd-F `madness is poor`).
- [ ] Unicode breaking if unicode is at the point of trim (16 bytes v/s 8 bytes I think) (try searching `hamlet`, then Cmd-F `a moiety`).
- [ ] Add feature for word-bound search.
