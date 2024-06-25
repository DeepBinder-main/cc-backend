<script>
    import { Icon, Container } from "sveltestrap";
    // import Home from 'lucide-svelte/icons/home';
    // FaCircleDot
    // import { FaCircleDot } from "svelte-icons-pack/fa";
</script>

<Container style="background-color: white; min-height: 800px;" class="mt-3 p-3">
    <div class="d-flex gap-3 flex-wrap">
        <div style="background-color: white; box-shadow: 0 0 10px grey;" class="card px-3 py-2">
            <div class="row border-3">
                <div class="col-2 mt-3">
                    <Icon class="fs-2 text-primary" name="dice-1-fill" />
                    <!-- <Home /> -->
                </div>
                <div class="col-8">
                    <h3>Logical Volumes</h3>
                    <p class="text-gray">12 Volumes</p>
                </div>
                <div class="col-2 mt-3">
                    <Icon class="fs-2" name="arrow-right" />
                </div>
            </div>
        </div>

        <div style="background-color: white; box-shadow: 0 0 10px grey;" class="card  px-3 py-2">
            <div class="row border-1">
                <div class="col-2 mt-3">
                    <Icon class="fs-2 text-success" name="box-fill" />
                </div>
                <div class="col-8">
                    <h3>Physical Volumes</h3>
                    <p class="text-gray">8 Volumes</p>
                </div>
                <div class="col-2 mt-3">
                    <Icon class="fs-2" name="arrow-right" />
                </div>
            </div>
        </div>
        
        <div style="background-color: white; box-shadow: 0 0 10px grey;" class="card  px-3 py-2">
            <div class="row border-1">
                <div class="col-2 mt-3">
                    <Icon class="fs-2 text-warning fw-bold" name="stack" />
                </div>
                <div class="col-8">
                    <h3>Volume Groups</h3>
                    <p class="text-gray">4 groups</p>
                </div>
                <div class="col-2 mt-3">
                    <Icon class="fs-2" name="arrow-right" />
                </div>
            </div>
            
        </div>

        <div style="background-color: white; box-shadow: 0 0 10px grey;" class="card  px-3 py-2">
            <div class="row border-1">
                <div class="col-2 mt-3">
                    <Icon class="fs-2 text-danger" name="diagram-3-fill" />
                    <!-- <Icon name="border" /> -->
                </div>
                <div class="col-8">
                    <h3>Storage Partitions</h3>
                    <p class="text-gray">12 Partitions</p>
                </div>
                <div class="col-2 mt-3">
                    <Icon class="fs-2" name="arrow-right" />
                </div>
            </div>
        </div>
    </div>
</Container>

<style>
    .card {
        width: 500px !important;
    }
</style>
