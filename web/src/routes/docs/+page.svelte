<svelte:head>
  <title>MLforge API Docs</title>
</svelte:head>

<div id="swagger-ui"></div>

<script>
  import { onMount } from 'svelte';
  import { browser } from '$app/environment';
  import 'swagger-ui-dist/swagger-ui.css';

  onMount(async () => {
    if (!browser) return;
    const { SwaggerUIBundle, SwaggerUIStandalonePreset } = await import('swagger-ui-dist');
    SwaggerUIBundle({
      url: '/openapi.yaml',
      dom_id: '#swagger-ui',
      deepLinking: true,
      persistAuthorization: true,
      presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
      layout: 'StandaloneLayout',
    });
  });
</script>
