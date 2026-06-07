<script>
  import { onMount } from 'svelte';
  import { currentStepIndex, currentPath } from './stores.js';
  import slugOrder from './slugOrder.js';

  export let stepName = '';
  export let location;

  import Step             from './Step.svelte';
  import Code             from './Code.svelte';
  import ExternalLink     from './ExternalLink.svelte';
  import ExternalButton   from './ExternalButton.svelte';
  import PrevNextStepButtons from './PrevNextStepButtons.svelte';
  import { serverHosts, chooseRandomHosts } from './serverHosts.js';

  onMount(() => {
    currentPath.set(window.location.pathname);

    if (stepName !== slugOrder[$currentStepIndex]) {
      if ($currentStepIndex === 0) {
        for (let i = 0; i < slugOrder.length; i++) {
          if (slugOrder[i] === stepName) {
            $currentStepIndex = i;
            break;
          }
        }
      }
    }

    window.scrollTo(0, 0);
  });

  const displayHosts = (hosts) => {
    const links = hosts.map(h =>
      `* ${h.name}: <a href="${h.url}" target="_blank" rel="nofollow noreferrer noopener">${h.url}</a> (${h.country})`
    );
    return links.join('<br />');
  };

  const numSuggestedProviders = 3;
  let viewAllHosts = false;
</script>

<style>
  .steps-wrap {
    display: flex;
    flex-direction: column;
  }

  p { margin-bottom: 14px; }

  .host-list {
    background: var(--bg-raised);
    border: 1px solid var(--border);
    border-radius: 8px;
    padding: 16px 20px;
    margin: 12px 0;
    font-size: 0.875rem;
    line-height: 1.9;
  }

  .view-all-btn {
    margin-top: 8px;
  }
</style>

<div class="steps-wrap">

  <Step slug="intro" title="Introduction &amp; Disclaimer">
    <p>
      Welcome to Run A Relay! Our goal is to make it as fast and easy
      as possible to set up a new
      <ExternalLink href="https://www.torproject.org/">Tor</ExternalLink>
      relay, with special emphasis on helping you choose a host based on
      <ExternalLink href="https://trac.torproject.org/projects/tor/wiki/doc/GoodBadISPs">
        crowdsourced information from The Tor Project
      </ExternalLink>.
    </p>
    <p>We've distilled it all down to actionable steps.</p>
  </Step>


  <Step slug="relay-or-donate" title="Run a Relay or Donate?">
    <p>
      If you have less than $10/month to spend on a server,
      consider donating to one of these projects instead —
      it will make the Tor network more efficient.
    </p>
    <p>
      <ExternalLink href="https://torservers.net/donate.html">TorServers.net</ExternalLink>
      &nbsp;·&nbsp;
      <ExternalLink href="https://noisetor.net/">Noisetor</ExternalLink>
      &nbsp;·&nbsp;
      <ExternalLink href="https://emeraldonion.org/donate/">Emerald Onion</ExternalLink>
    </p>
  </Step>


  <Step slug="choosing-a-host" title="Choosing a Host">
    <p>
      Server providers to consider — refresh for another
      {numSuggestedProviders} randomly-chosen providers from our curated list:
    </p>

    {#if viewAllHosts}
      <div class="host-list">{@html displayHosts(serverHosts)}</div>
    {:else}
      <div class="host-list">{@html displayHosts(chooseRandomHosts(numSuggestedProviders))}</div>
      <div class="view-all-btn">
        <button on:click={() => viewAllHosts = true}>View all providers</button>
      </div>
    {/if}
  </Step>


  <Step slug="install-tor-daemon" title="Install Tor &amp; Configure Your Relay">
    <p>
      Use the impressively awesome
      <ExternalLink href="https://github.com/nusenu/ansible-relayor">ansible-relayor</ExternalLink>
      to set up your relay on Debian, Ubuntu, FreeBSD, or OpenBSD.
    </p>
    <p>Non-Debian is preferred for network diversity.</p>
    <p>
      Or set up without Ansible:
      <ExternalLink href="https://community.torproject.org/relay/setup/guard/">
        community.torproject.org/relay/setup/guard/
      </ExternalLink>
    </p>
  </Step>


  <Step slug="thank-you" title="You're All Set!">
    <p>
      <em>Thank you so much</em> for taking action to
      make the Tor network faster and more scalable.
    </p>
    <p>
      Each new relay can potentially help thousands of people protect
      their privacy every single day.
    </p>
    <p style="margin-top: 20px;">
      <ExternalButton primary href="https://twitter.com/intent/tweet?text=I%20just%20spun%20up%20a%20new%20@TorProject%20relay!%20Simple%20instructions%20for%20running%20your%20own:%20https://RunARelay.org">
        Share on Twitter
      </ExternalButton>
    </p>
    <p style="margin-top: 24px;">
      For more info: <ExternalLink href="https://trac.torproject.org/projects/tor/wiki/TorRelayGuide">Tor Relay Guide ↗</ExternalLink>
    </p>
  </Step>

  <PrevNextStepButtons />
</div>
