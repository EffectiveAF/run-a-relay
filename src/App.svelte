<script>
  import { Router, Route } from 'svelte-routing';

  import Nav from './Nav.svelte';
  import Home from './Home.svelte';
  import About from './About.svelte';
</script>

<Router>
  <Nav />
  <div>
    <Route path="/"><Home /></Route>
    <Route path="/about"><About /></Route>
  </div>
</Router>
