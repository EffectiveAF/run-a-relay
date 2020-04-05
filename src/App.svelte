<script>
  import { Router, Route } from 'svelte-routing';

  import Nav from './Nav.svelte';
  import Home from './Home.svelte';
</script>

<Router>
  <Nav />
  <div>
    <Route path="/"><Home /></Route>
  </div>
</Router>
