<script lang="ts">
    let {
        title,
        description
    }: {
        title: string;
        description: string;
    } = $props();
</script>

<div>
    <h1 class="h3 mb-1">{title}</h1>
    <p class="text-body-secondary mb-0">{description}</p>
</div>