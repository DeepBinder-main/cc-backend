<script >
    export let name;
    export let progress;
    import { Row, Col, Card, CardBody, CardTitle, Progress } from "sveltestrap";
</script>

<Col>
    <CardTitle>{name}</CardTitle>
    <Progress value={progress}>{progress}%</Progress>
</Col>
