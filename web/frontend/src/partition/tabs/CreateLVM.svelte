<script>
    import { Input, Button} from 'sveltestrap';
    
    let formData = {
        volumeName : "",
        mountPoint : "",
        fileType : "",
        size : "",
    }

    function onSubmit(e) {
        console.log(formData);
        // formData = {
        //     volumeName : "",
        //     mountPoint : "",
        //     fileType : "",
        //     size : "",
        // } 
    }
</script>

<form on:submit|preventDefault={onSubmit}>
    <h1 style="text-align: center;">Create LVM</h1>
    <div class="d-flex  justify-content-between">
        <div class="ms-5 row gap-3">
            <Input bind:value="{formData.volumeName}" placeholder="Enter Volume Name" type="text" />
            <Input bind:value="{formData.size}" placeholder="Enter Size" type="text" />
        </div>
        <div class="ms-5 row gap-3">
            <Input bind:value="{formData.mountPoint}" placeholder="Enter Mount Point" type="text" />
            <Input bind:value="{formData.fileType}" placeholder="Enter File System type" type="text" />
        </div>
    
    </div>
    <Button class="ms-5 mt-3" type="submit">Submit</Button>
</form>