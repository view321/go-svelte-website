<!-- frontend/src/App.svelte -->
<script>
	import { onMount } from 'svelte';

	let message = 'Loading...';

	onMount(async () => {
		try {
			const response = await fetch('/api/message'); // Fetch from Go backend
			const data = await response.json();
			message = data.text;
		} catch (error) {
			console.error('Error fetching message:', error);
			message = 'Error loading message.';
		}
	});
</script>

<main>
	<h1>Svelte and Go</h1>
	<p>{message}</p>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>