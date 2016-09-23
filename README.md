#twitcheck
twitcheck is a quick and simple command-line tool to let you checkin on a Twitter user. And tell you what word they have been using most frequently.

##Installing
We suggest using the lovely go get tool to quickly add it to your path.
```
go get -u github.com/mvryan/twitcheck
```
Additionally twitcheck searches the current directory for a configuration file.

config.json
```json
{
	"ConsumerKey":"YOUR-KEY-HERE",
	"ConsumerSecret":"YOUR-SECRET-HERE"
}

```

##Usage
Simply supply twitcheck with a user name and it will pull recent information about them and provide information about the calls to the API.
```
Imin:~ ryan$ twitcheck twitterdev
-------------------Profile-------------------
Name:                 TwitterDev
Screen Name:          TwitterDev
Description:          Developer and Platform Relations @Twitter. We are developer advocates. We can't answer all your questions, but we listen to all of them!
Location:             Internet
Tweets:               2642
Followers:            431542
Following:            1534

---------------------Info--------------------
Status:               200 OK
Time Elapsed:         0.331737s
Rate limit:           181
Rate limit remaining: 180
Rate limit reset:     2016-09-22 21:36:47 -0400 EDT

-------------------Tweets-------------------
1 : 👏 👊 @Comprenditech for winning our #Promote Ads API challenge 🏆 https://t.co/oMvfE7qbuE
	By, TwitterDev at 2016-09-22 00:45:31 -0400 EDT

2 : RT @AdsAPI: Less than 24 hours til #Promote Demo Day! Excited to see our top 12 share what they built. Stay tuned to learn who takes home t…
	By, TwitterDev at 2016-09-21 15:30:47 -0400 EDT

3 : NYC developers: Are you passionate about Twitter's 🛠 &amp; advocating for Twitter ❤️ locally? Learn more next Tuesday. https://t.co/1YB9VfTonK
	By, TwitterDev at 2016-09-21 15:17:06 -0400 EDT

4 : RT @wocintechchat: Inaugural Meeting of @TwitterDev Community @TwitterNYC on Tuesday 9/27. #WOCinTech You can RSVP here --&gt; https://t.co/E…
	By, TwitterDev at 2016-09-21 15:02:10 -0400 EDT

5 : RT @andypiper: Have you updated your @twitterapi app for the new attachment_urls yet? https://t.co/J0NmSMHkHP - questions? ask via https://…
	By, TwitterDev at 2016-09-20 19:20:23 -0400 EDT

Most Common Word Tweeted: rt
---------------------Info--------------------
Status:               200 OK
Time Elapsed:         0.159700s
Rate limit:           300
Rate limit remaining: 299
Rate limit reset:     2016-09-22 21:36:47 -0400 EDT
```
