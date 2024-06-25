<script>
    import { Bar } from "svelte-chartjs";
    import {Card, Row} from 'sveltestrap';

    import {
        Chart,
        Title,
        Tooltip,
        Legend,
        BarElement,
        CategoryScale,
        LinearScale,
    } from "chart.js";

    Chart.register(
        Title,
        Tooltip,
        Legend,
        BarElement,
        CategoryScale,
        LinearScale,
    );

    let chartData = {
        labels: ['January', 'February', 'March', 'April', 'May'],
        datasets: [{
            label: 'Dataset 1',
            data: [10, 20, 30, 40, 50],
            backgroundColor: 'rgba(255, 99, 132, 0.5)'
        }, {
            label: 'Dataset 2',
            data: [15, 25, 35, 45, 55],
            backgroundColor: 'rgba(54, 162, 235, 0.5)'
        }]
    };

    const data = {
        labels: ['January', 'February', 'March', 'April', 'May'],
        datasets: [
            {
                label: "Used",
                backgroundColor: "red",
                data: [10, 20, 30, 40, 50],
                barThickness: 15,
            },
            {
                label: "Free",
                backgroundColor: "blue",
                data: [15, 25, 35, 45, 55],
                barThickness: 15,
            },
        ],
    };

    const options = {
        responsive: true,
        indexAxis: "y",
        scales: {
            x: {
                stacked: true,
            },
            y: {
                stacked: true,
            },
        },
    };
</script>

<div class="row mt-4">
    <div class="col">
        <select name="" id="" class="form-select">
            <option value="">Logical Group</option>
            <option value="">Ubuntu-vg</option>
            <option value="">MAC-vg</option>
            <option value="">Windows-vg</option>
        </select>
    </div>
    <div class="col">
        <select name="" id="" class="form-select">
            <option value="">select file system type</option>
            <option value="">ext2</option>
            <option value="">ext3</option>
            <option value="">ext4</option>
        </select>
    </div>
    <div class="col">
        <input type="text" placeholder="Enter size of volume" class="form-control">
    </div>
</div>

<div class="row mt-4">
    <div class="col-10">
        <input type="text" placeholder="Enter logical volume name" class="form-control">
    </div>
    <div class="col-2">
        <div class="d-grid">
            <button class="btn btn-dark">create LVM</button>
        </div>
    </div>
</div>

<Row cols={2} class="mt-4 ps-3">
    <Card style="width: 49%;" class="p-3">
        <h3>Used Space</h3>
        <Bar style="height: 70px; " {data} options={options} />
    </Card>
    
    <Card style="width: 49%;" class="ms-3 p-3">
        <table class="table table-hover">
            <thead class="table-dark">
                <tr>
                    <th>Logical Volume</th>
                    <th>Logical Group</th>
                    <th>Size(GB)</th>
                    <th>File System</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>lv01</td>
                    <td>Ubuntu-vg</td>
                    <td>50</td>
                    <td>ext4</td>
                    <td>
                        <button class="btn btn-danger">Delete</button>
                    </td>
                </tr>
                <tr>
                    <td>lv02</td>
                    <td>Ubuntu-vg</td>
                    <td>20</td>
                    <td>ext4</td>
                    <td>
                        <button class="btn btn-danger">Delete</button>
                    </td>
                </tr>
                <tr>
                    <td>lv03</td>
                    <td>Ubuntu-vg</td>
                    <td>10</td>
                    <td>ext4</td>
                    <td>
                        <button class="btn btn-danger">Delete</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </Card>
</Row>



