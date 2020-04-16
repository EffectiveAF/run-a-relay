<script>
  import ExternalLink from './ExternalLink.svelte';

  export let location;  // implicitly populated by router
</script>

<style>
</style>

<div class="static-content">
  <h3>Who Built <em>Run A Relay</em>?</h3>

  <ul>
    <li>- Steve Phillips / <ExternalLink href="https://github.com/elimisteve">@elimisteve</ExternalLink>, privacy activist and Founder of <ExternalLink href="https://effective.af/">EffectiveAF</ExternalLink></li>
    <ul>
      <li>- Lead developer and hoster of <ExternalLink href="https://leapchat.org">LeapChat</ExternalLink>: end-to-end encrypted, ephemeral chat in the browser</li>
    </ul>
    <li>- Tim Sullivan / <ExternalLink href="https://github.com/timsully">@timsully</ExternalLink>, Design Technologist at Splunk</li>
  </ul>

  <p>
    We love Tor!
  </p>
</div>
