<script lang="ts">
    import Card from "$lib/components/common/Card.svelte";
    import MLforgeSideBar from "$lib/components/common/MLforgeSideBar.svelte";
    import MLforgeSideBarItem from "$lib/components/common/MLforgeSideBarItem.svelte";
    import Title from "$lib/components/common/Title.svelte";
    import NamedCard from "$lib/components/common/NamedCard.svelte";
    
    const runs = Array.from({ length: 30 }, (_, i) => ({
        id: i + 1,
        name: `Run ${i + 1}`,
        accuracy: (0.80 + i * 0.003).toFixed(3),
        loss: (0.50 - i * 0.01).toFixed(3)
    }));
</script>

<div class="d-flex vh-100 overflow-hidden">

    <MLforgeSideBar>
        <MLforgeSideBarItem path="/experiment1" active>
            Experiment 1
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment2">
            Experiment 2
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment3">
            Experiment 3
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment4">
            Experiment 4
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment5">
            Experiment 5
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment6">
            Experiment 6
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment7">
            Experiment 7
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment8">
            Experiment 8
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment9">
            Experiment 9
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment10">
            Experiment 10
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
        <MLforgeSideBarItem path="/experiment11">
            Experiment 11
        </MLforgeSideBarItem>
    </MLforgeSideBar>

    <main class="flex-grow-1 overflow-auto p-4">
        <div class="container-fluid">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <Title 
                    title="Create New Experiment" 
                    description="Create a new experiment to start tracking your training runs."
                />
                <button class="btn btn-primary">Create</button>
            </div>
            <div class="row g-3">
                <NamedCard
                    title="Experiment Name"
                    description="Experiment Name"
                >
                    <input 
                        type="text" 
                        class="form-control" 
                        id="experimentFormControlInput" 
                        placeholder="Experiment 1"
                    />
                </NamedCard>
                <NamedCard
                    title="Experiment Description"
                    description="Experiment Description"
                >
                    <textarea 
                        class="form-control" 
                        id="experimentFormControlTextarea" 
                        rows="3"
                    ></textarea>
                </NamedCard>
            </div>
        </div>
    </main>
</div>
