<script>
  import ExternalLink from './ExternalLink.svelte';

  export let location;  // implicitly populated by router
</script>

<div class="static-content">
  <h3>Terms of Service</h3>

  <p>
    EffectiveAF is a proud supporter of the
    <ExternalLink href="https://ethicalsource.dev/">Ethical Source</ExternalLink>
    movement.  Run-A-Relay is released under an
    Ethical Source license, namely The Hippocratic License 2.1 (see
    below).
  </p>

  <p>
    You may do as you wish with this website and
    <ExternalLink href="https://github.com/EffectiveAF/run-a-relay">its software</ExternalLink>
    provided that you do not somehow use it to
    violate human rights as defined by the UN Universal Declaration of
    Human Rights;
    see <ExternalLink href="https://github.com/EffectiveAF/run-a-relay/blob/master/LICENSE.md">LICENSE.md</ExternalLink>
    or <ExternalLink href="https://firstdonoharm.dev/">The Hippocratic
    License 2.1</ExternalLink> for details.
  </p>
</div>
