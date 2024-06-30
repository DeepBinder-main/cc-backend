<script>
    import { Card, Input } from "sveltestrap";
</script>


<h1 class="text-center">Deploy New Agent</h1>
<div class="d-flex justify-content-center mt-3 ">
    <div style="width: 85%;" >
        <h3>Select Package to download</h3>
        <div class="d-flex justify-content-center gap-3">
            <Card style="width: 350px;" class=" p-3">
                <h2 class="mb-0 text-center">Linux</h2>
                <hr>
                <div class="row">
                    <div class="col">
                        <Input type="radio" label="DEB arm64" name="radio1" />
                        <Input type="radio" label="RPM arm64" name="radio1" />
                    </div>
                    
                    <div class="col">
                        <Input type="radio" label="DEB arch64" name="radio1" />
                        <Input type="radio" label="RPM arch64" name="radio1" />
                    </div>
                </div>
            </Card>
            <Card style="width: 350px;" class=" p-3">
                <h2 class="mb-0 text-center">Windows</h2>
                <span class="text-danger">*currently under development</span>
                <hr>
                <Input type="radio" label="MSI 32/64 bit" name="radio1" />
            </Card>
            <!-- <Card style="width: 350px;" class=" p-3">
                <h2 class="mb-0 text-center">mac OS</h2>
                <hr>
                <Input type="radio" label="intel" name="radio1" />
                <Input type="radio" label="Apple silicon" name="radio1" />
            </Card> -->
        </div>
    </div>
</div>

<div class="d-flex justify-content-center mt-3">
    <div style="width: 85%; margin-top: 20px;" >
        <h3>Assign Server Address</h3>
        Lorem ipsum, dolor sit amet consectetur adipisicing elit. Eveniet iste exercitationem tempora iusto odit quam iure corporis harum debitis laboriosam!
    
        <Input placeholder="Assign Server Address" />
    </div>
</div>

<div class="d-flex justify-content-center mt-3">
    <div style="width: 85%; margin-top: 20px;" >
        <h3>Optional Settings</h3>
        Lorem ipsum dolor sit, amet consectetur adipisicing elit. Nostrum, laudantium. Similique minus temporibus quos placeat.
        <form>
            <div class="form-group mt-3">
                <label for="an" class="form-label">Assign Agent Name</label>
                <input class="form-control" type="text" name="name" id="an">
            </div>
            <div class="form-group mt-2">
                <label for="grp" class="form-label">Select group</label>
                <select name="group" id="grp" class="form-select" >
                    <option value="1">Windows</option>
                    <option value="1">Windows</option>
                    <option value="1">Windows</option>
                    <option value="1">Windows</option>
                </select>
            </div>
        </form>
    </div>
</div>

<div class="d-flex justify-content-center mt-3 ">
    <div style="width: 85%; margin-top: 20px;" >
        <h3>Run The following</h3>
        <textarea name="cmd" id="c,d" class="form-control">
            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Aut reprehenderit dolorum iste neque dicta, rem ipsam nostrum incidunt aspernatur laborum.
        </textarea>
    </div>
</div>

<div class="d-flex justify-content-center mt-3 mb-5 pb-3">
    <div style="width: 85%; margin-top: 20px;" >
        <h3>Start the agent</h3>
        <input name="cmd" id="cmd" value="NET START Lorem" class="form-control" />
    </div>
</div>