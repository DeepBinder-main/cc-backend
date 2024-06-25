<script>
    import {Input, Container, Button} from 'sveltestrap'
   
</script>

<Container>
    <div class="mt-3">
        <form>
            <div class="row mt-3">
                <label class="col-1" for="name">Username</label>
                <div class="col-10">
                    <Input type="text" id="name" />
                </div>
            </div>
            <div class="row mt-3">
                <label class="col-1" for="email">Email</label>
                <div class="col-10">
                    <Input type="email" id="email" />
                </div>
            </div>
            <div class="row mt-3">
                <label class="col-1" for="pass1">Password</label>
                <div class="col-10">
                    <Input id="pass1" />
                </div>
            </div>
            <div class="row mt-3">
                <label class="col-1" for="pass2"> Reenter Password</label>
                <div class="col-10">
                    <Input id="pass2" />
                </div>
            </div>
        </form>
    </div>
    <hr >
    
    <div class="mt-3">
        <h4>Groups</h4>
        
        <Input type="checkbox" label="Group1" />
        <Input type="checkbox" label="Group2" />
        <Input type="checkbox" label="Group3" />
    </div>
    <hr >
    
    <div class="mt-3">
        <h4>Account Restrictions</h4>
        <Input type="checkbox" label="Login Disable" />
    </div>
    <hr>
    
    <div class="mt-3">
        <h4>Mount Logical Volume</h4>
        <select class="form-select" name="vol" id="vol">
            <option value="">select volume</option>
            <option value="">Volume1</option>
            <option value="">Volume2</option>
            <option value="">Volume3</option>
        </select>
    </div>
    <hr>
    
    <div class="mt-3">
        <h4>Public Key</h4>
        <textarea name="public-key" class="form-control"></textarea>
    </div>
    <hr>

    <div class="d-flex justify-content-center">
        <Button color="dark">save</Button>
        <Button color="dark" class="ms-2">clone</Button>
        <Button color="dark" class="ms-2">cancel</Button>
        <Button color="danger" class="ms-2">delete</Button>
    </div>
</Container>
