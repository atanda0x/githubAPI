GitHub provides a well-written REST API to consume from the users. It opens up the data about users, repositories, repository statistics, and so on, to the clients through the API. The current stable version is v3. The API documentation can be found at 

```
https://developer.github.com/v3/. 

```

The root endpoint of the API is:

```
curl https://api.github.com

```

The other API will be added to this base API. Now see how to make a few queries and get data related to various elements. For the unauthenticated user, the rate limit is 60/hour, whereas for clients who are passing client_id (which one can get from the GitHub account), it is 5,000/hour.

If you have a GitHub account (if not, it is recommended you create one), you can find accesstokens in the Your Profile | Personal Access Tokens area or by visiting.

```
https://github.com/settings/tokens. 

```
Create a new access token using the Generate new token button. Itasks for various permissions for types for the resource. Tick all of them. A new string will be generated. Save it to some private place. The token we have generated can be used to access the GitHub API (for longer rate limits).

The next step is to save that access token to an environment variable, GITHUB_TOKEN. To do
that, open your ~/.profile or ~/.bashrc file and add this as the last line:
```
export GITHUB_TOKEN=YOUR_GITHUB_ACCESS_TOKEN

```

YOUR_GITHUB_ACCESS_TOKEN is what was generated and saved previously from the
GitHub account.

