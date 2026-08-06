<script lang="ts">
	import 'swagger-ui-dist/swagger-ui.css';
	import { onMount } from 'svelte';

	onMount(async () => {
		const [{ default: SwaggerUIBundle }, { default: SwaggerUIStandalonePreset }] = await Promise.all([
			import('swagger-ui-dist/swagger-ui-bundle.js'),
			import('swagger-ui-dist/swagger-ui-standalone-preset.js')
		]);

		SwaggerUIBundle({
			url: '/openapi.yaml',
			dom_id: '#swagger-ui',
			deepLinking: true,
			presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
			layout: 'StandaloneLayout'
		});
	});
</script>

<div id="swagger-ui"></div>
