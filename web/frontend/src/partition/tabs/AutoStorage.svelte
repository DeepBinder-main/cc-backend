<script>
    import { Button, Container } from "sveltestrap";
</script>

<Container>
    <div class="row mt-5">
        <div class="col-12 col-md-6 mb-3">
            <h3>Auto-Extend Settings</h3>
            <div>
                <form class="row">
                    <div class="col-4">
                        <div class="form-label">Username</div>
                        <input
                            type="text"
                            class="form-control"
                            placeholder="Enter username"
                        />
                    </div>
                    <div class="col-4">
                        <div class="form-label">Alert type</div>
                        <select class="form-select">
                            <option>select alert type</option>
                            <option>slack</option>
                            <option>SMTP</option>
                        </select>
                    </div>
                    <div class="col-4 mt-4">
                        <Button color="dark">Submit</Button>
                    </div>
                </form>
            </div>
        </div>
        
        <div class="col-12 col-md-6">
            <h3>Auto-Reduce Settings</h3>
            <div>
                <form class="row">
                    <div class="col-4">
                        <div class="form-label">Threshold</div>
                        <input
                            type="text"
                            class="form-control"
                            placeholder="Enter username"
                        />
                    </div>
                    <div class="col-4">
                        <div class="form-label">Select Volume</div>
                        <select class="form-select">
                            <option>select select volume</option>
                            <option>volume1</option>
                            <option selected>volume2</option>
                            <option>volume3</option>
                        </select>
                    </div>
                    <div class="col-4 mt-4">
                        <Button color="dark">Submit</Button>
                    </div>
                </form>
            </div>
        </div>
    </div>
    
    <div class="row mt-5">
        <h3>Users</h3>
        <table class="table table-hover">
            <thead class="table-dark">
                <tr class="">
                    <th> Name </th>
                    <th> Used Storage </th>
                    <th> Total Storage </th>
                    <th> Auto-Extend </th>
                    <th> Auto-Reduce </th>
                    <th> Action </th>
                </tr>
            </thead>
            <tbody class="">
                <tr>
                    <td>John Doe</td>
                    <td>80GB</td>
                    <td>100GB</td>
                    <td>90% (vol1)</td>
                    <td>70% (vol2)</td>
                    <td>
                        <Button outline >Auto Extend</Button>
                        <Button outline >Auto Reduce</Button>
                    </td>
                </tr>
                <tr class="">
                    <td>Jane Smith</td>
                    <td>60GB</td>
                    <td>100GB</td>
                    <td>85% (vol3)</td>
                    <td>75% (vol4)</td>
                    <td>
                        <Button outline >Auto Extend</Button>
                        <Button outline >Auto Reduce</Button>
                    </td>
                </tr>
                <tr class="">
                    <td>Bob Johnson</td>
                    <td>90GB</td>
                    <td>100GB</td>
                    <td>95% (vol5)</td>
                    <td>80% (vol6)</td>
                    <td>
                        <Button outline >Auto Extend</Button>
                        <Button outline >Auto Reduce</Button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</Container>
