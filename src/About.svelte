<script>
  import ExternalLink from './ExternalLink.svelte';

  export let location;  // implicitly populated by router
</script>

<style>
  .about {
    padding: 40px;
    font-size: 18px;
  }
  ul {
    margin-left: 30px;
  }
  li, p {
    margin-top: 10px;
  }
</style>

<div class="about">
  <h2>Who Built <em>Run A Relay</em>?</h2>

  <ul>
    <li>- Steve Phillips / <ExternalLink href="https://github.com/elimisteve">@elimisteve</ExternalLink>, privacy activist and Founder of <ExternalLink href="https://effective.af/">EffectiveAF</ExternalLink></li>
    <ul>
      <li>- Lead developer and hoster of <ExternalLink href="https://leapchat.org">LeapChat</ExternalLink>: end-to-end encrypted, ephemeral chat in the browser</li>
    </ul>
    <li>- Tim Sullivan / <ExternalLink href="https://github.com/timsully">@timsully</ExternalLink>, UI/human factors engineer at Splunk</li>
  </ul>

  <p>
    We love Tor!
  </p>
</div>
