<script>
  import { onMount } from 'svelte';
  import { currentPath } from './stores.js';
  import ExternalLink from './ExternalLink.svelte';

  export let location;
  onMount(() => currentPath.set('/about'));
</script>

<style>
  .page-header {
    padding-bottom: 28px;
    margin-bottom: 28px;
    border-bottom: 1px solid var(--border);
  }

  .page-header p {
    color: var(--text-2);
    margin: 6px 0 0;
    font-size: 0.9375rem;
  }

  .team-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 14px;
    margin-top: 16px;
  }

  @media (max-width: 480px) {
    .team-grid { grid-template-columns: 1fr; }
  }

  .team-card {
    background: var(--bg-raised);
    border: 1px solid var(--border);
    border-radius: 10px;
    overflow: hidden;
    transition: border-color 0.12s;
  }

  .team-card:hover { border-color: var(--border-strong); }

  .team-card img {
    width: 100%;
    aspect-ratio: 3/4;
    object-fit: cover;
    object-position: top center;
    display: block;
  }

  .team-info {
    padding: 14px 16px;
  }

  .team-name {
    font-size: 0.9rem;
    font-weight: 600;
    color: var(--text);
    margin-bottom: 3px;
  }

  .team-role {
    font-size: 0.8rem;
    color: var(--text-2);
    line-height: 1.4;
  }
</style>

<div class="page-content">

  <div class="page-header">
    <h1>Creating a better future, together.</h1>
    <p>A free, open-source guide to running Tor relays and protecting internet privacy.</p>
  </div>

  <h2>Who we are</h2>
  <p>
    We're a small team of engineers and designers who believe the internet should
    remain free and private for everyone. Run A Relay was built to lower the barrier
    to contributing to the Tor network.
  </p>

  <h2>Our mission</h2>
  <p>
    To make it as easy as possible for anyone with a spare $10/month and 30 minutes
    to spin up a Tor relay — strengthening the network one node at a time.
  </p>

  <h2>The team</h2>
  <div class="team-grid">
    <div class="team-card">
      <img src="../img/steve.png" alt="Steve Phillips" />
      <div class="team-info">
        <div class="team-name">Steve Phillips</div>
        <div class="team-role">
          Founder, <ExternalLink href="https://effective.af/">Effective.AF</ExternalLink>
        </div>
      </div>
    </div>

    <div class="team-card">
      <img src="../img/tim.png" alt="Tim Sullivan" />
      <div class="team-info">
        <div class="team-name">Tim Sullivan</div>
        <div class="team-role">Design Technologist</div>
      </div>
    </div>
  </div>

</div>
