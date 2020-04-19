<script>
  import { onMount } from 'svelte';
  import { currentStepIndex } from './stores.js';
  import slugOrder from './slugOrder.js';

  export let stepName = '';
  export let location;  // implicitly populated by router

  import Step from './Step.svelte';
  import Code from './Code.svelte';
  import ExternalLink from './ExternalLink.svelte';
  import ExternalButton from './ExternalButton.svelte';
  import PrevNextStepButtons from './PrevNextStepButtons.svelte';
  import { serverHosts, chooseRandomHosts } from './serverHosts.js';

  onMount(async () => {
    if (stepName !== slugOrder[$currentStepIndex]) {
      if ($currentStepIndex === 0) {
        // User was linked to /step/some-specific-step, so take them there
        for (let i = 0; i < slugOrder.length; i++) {
          if (slugOrder[i] === stepName) {
            $currentStepIndex = i;
            break;
          }
        }
      } else {
        console.log('Something unexpected occurred');
      }
    }

    // Useful on mobile when coming from homepage
    window.scrollTo(0, 0);
  })

  const displayHosts = (hosts) => {
    const links = hosts.map(h => `* ${h.name}: <a href="${h.url}" target="_blank" rel="nofollow noreferrer noopener">${h.url}</a> (${h.country})`);
    return links.join('<br />');
  }

  const numSuggestedProviders = 3;
</script>

<style>
  .all-steps {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: left;
  }

  p {
    margin-bottom: 16px;
  }
</style>

<div class="all-steps">
  <Step slug="intro" title="Introduction and Disclaimer">
    <p>
      Welcome to Run A Relay! Our goal is to make it as fast and easy
      as possible to stand up a new
      <ExternalLink href="https://www.torproject.org/">Tor</ExternalLink>
      relay, with special emphasis placed on helping you
      choose a host/server provider based on
      <ExternalLink href="https://trac.torproject.org/projects/tor/wiki/doc/GoodBadISPs">
        crowdsourced information from The Tor Project's wiki
      </ExternalLink>.
    </p>

    <p>
      We've distilled it all down to actionable steps!
    </p>
  </Step>


  <Step slug="relay-or-donate" title="Run or Donate?">
    <p>
      If you have less than $10/month that you can spend on a server
      to run your relay, consider donating to one of these fine
      projects instead!  This will make the Tor network more
      efficient.
    </p>

    <p>
      <ExternalLink href="https://torservers.net/donate.html">TorServers.net</ExternalLink>
      &nbsp;|&nbsp;
      <ExternalLink href="https://noisetor.net/">Noisetor</ExternalLink>
      &nbsp;|&nbsp;
      <ExternalLink href="https://emeraldonion.org/donate/">Emerald Onion</ExternalLink>
    </p>
  </Step>


  <Step slug="choosing-a-host" title="Choosing A Host">
    <p>
      Server providers to consider (refresh for another
      {numSuggestedProviders} randomly-chosen providers from our
      curated list!):
    </p>

    <p>
      {@html displayHosts(chooseRandomHosts(numSuggestedProviders))}
    </p>
  </Step>


  <Step slug="install-tor-daemon" title="Install Tor Daemon & Configure Relay">
    <p>
      Use the impressively
      awesome <ExternalLink href="https://github.com/nusenu/ansible-relayor">https://github.com/nusenu/ansible-relayor</ExternalLink>
      to stand up your relay on Debian, Ubuntu, FreeBSD, or OpenBSD!
    </p>

    <p>
      (Non-Debian preferred!  Because diversity.)
    </p>
  </Step>


  <!--
    NOTE: When adding a new <Step>, remember to add its slug to ./slugOrder.js!
  -->


  <Step slug="thank-you" title="Thank You!">
    <p>
      You're all set!  <em>Thank you so much</em> for taking action to
      make the Tor network faster and more scalable!
    </p>

    <p>
      Each new relay can potentially help thousands of people protect
      their privacy every single day.  The more the merrier!
    </p>

    <p style="margin-top: 25px; margin-bottom: 25px;">
      <ExternalButton primary href="https://twitter.com/intent/tweet?text=I%20just%20spun%20up%20a%20new%20@TorProject%20relay!%20Simple%20instructions%20for%20running%20your%20own:%20https://RunARelay.org">Brag on Twitter</ExternalButton>
    </p>

    <p>
      For more information, check out the official Tor Relay Guide:
      <ExternalLink href="https://trac.torproject.org/projects/tor/wiki/TorRelayGuide">https://trac.torproject.org/projects/tor/wiki/TorRelayGuide</ExternalLink>
      .
    </p>
  </Step>

  <PrevNextStepButtons />
</div>
