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
  })
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

    <p style="margin-top: 25px;">
      <ExternalButton primary href="https://twitter.com/intent/tweet?text=I%20just%20spun%20up%20a%20new%20@TorProject%20relay!%20Simple%20instructions%20for%20running%20your%20own:%20https://RunARelay.org">Tweet about it</ExternalButton>
    </p>
  </Step>

  <PrevNextStepButtons />
</div>
