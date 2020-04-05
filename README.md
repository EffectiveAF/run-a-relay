## Run A Relay

Run A Relay is a web app that gives you simple instructions for
running a relay for the [Tor](https://www.torproject.org) anonymity
network!

Our goal is to dramatically increase the capacity of the Tor network.
This will make browsing the web through Tor even faster, and enable
Tor to protect even more people --
[millions per day and counting!](https://metrics.torproject.org/userstats-relay-country.html)


### FAQ

#### How does "Run A Relay" make it simpler to, you know, run a relay?

(1) By helping you select which server hosting providers to consider
(based on a number of factors; see below), and

(2) By telling you which command(s) to run to stand up your relay.

TODO(elimisteve): Say something about helping them choose which OS to
use if we indeed add this feature.


#### How does "Run A Relay" decide which hosting providers to consider?

Based on a number of factors:

(1) The information crowdsourced by the Tor community onto
[Tor's wiki](https://trac.torproject.org/projects/tor/wiki/doc/GoodBadISPs),
and

(2) Which providers are already commonly used to run relays as
measured by [Tor Metrics](https://metrics.torproject.org/).


## Running "Run A Relay" Locally

```
git clone http://github.com/EffectiveAF/run-a-relay
cd run-a-relay
npm install
npm run dev
```

Then just visit [localhost:5000](http://localhost:5000/)!


## Attribution

- `public/img/External_link_font_awesome.svg`
  - Font Awesome by Dave Gandy - https://fontawesome.github.com/Font-Awesome / CC BY-SA


## Donate to Tor

https://donate.torproject.org/


## Who Built "Run A Relay"?

* Steve Phillips / [@elimisteve](https://github.com/elimisteve), privacy activist and Founder of [EffectiveAF](https://effective.af/)
  * Lead developer and hoster of [LeapChat](https://leapchat.org): end-to-end encrypted, ephemeral chat in the browser

* Tim Sullivan / [@timsully](https://github.com/timsully), UI/human factors engineer at Splunk

We :heart: Tor!
