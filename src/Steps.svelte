<script>
  import { currentStepIndex } from './stores.js';
  import slugOrder from './slugOrder.js';

  export let stepName = '';

  import Step from './Step.svelte';
  import Code from './Code.svelte';
  import ExternalLink from './ExternalLink.svelte';
  import PrevNextStepButtons from './PrevNextStepButtons.svelte';
</script>

<style>
  .all-steps {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: left;
  }
</style>

<div class="all-steps">
  <Step slug="intro" title="Introduction and Disclaimer">
    <p>
      The following recommendations are based on crowdsourced information whose veracity we cannot guarantee.
    </p>
  </Step>


  <Step slug="relay-or-donate" title="Run v. Donate">
    <p>
      If you have less than $10/month you can spend on a server running a relay, consider donating to one of these fine projects instead!  That will make the network more efficient.
    </p>

    <div>
      <ExternalLink href="https://torservers.net/donate.html">TorServers.net</ExternalLink>
      &nbsp;|&nbsp;
      <ExternalLink href="https://noisetor.net/">Noisetor</ExternalLink>
      &nbsp;|&nbsp;
      <ExternalLink href="https://emeraldonion.org/">Emerald Onion</ExternalLink>
    </div>
  </Step>


  <Step slug="choosing-a-host" title="Choosing A Host">
    <p>
      Hosting providers to consider: ...
    </p>
  </Step>


  <Step slug="install-tor-daemon" title="Install Tor Daemon">
    <p>
      Commands to run:
    </p>

    <Code line="sudo apt-get install tor" />
  </Step>


  <!--
    NOTE: When adding a new <Step>, remember to add its slug to ./slugOrder.js!
  -->


  <Step slug="thank-you" title="Thank You!">
    <p>
      Thank you so much for taking action to make the Tor network faster and more scalable!
    </p>
  </Step>

  <PrevNextStepButtons />
</div>
