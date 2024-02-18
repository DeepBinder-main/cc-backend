<script lang="ts">
    import { Icon, TabContent, TabPane } from "sveltestrap";
    import {LayoutDashboard, DatabaseZap, Activity , Wrench ,Workflow} from "lucide-svelte";
    import Dashboard from "./Content/Dashboard.svelte";
</script>

<TabContent >
    <TabPane tabId="dashboard" active class="mt-4">
        <span slot="tab" class="small">
            <!-- <Icon name="gear" /> -->
            <LayoutDashboard />
            Dashboard 
        </span>
        <Dashboard />
    </TabPane>
    <TabPane tabId="createlv" class="mt-4">
        <span slot="tab" class="small">
            <DatabaseZap />
            Create LV
             <!-- <Icon name="hand-thumbs-up" /> -->
        </span>
        <Dashboard />
        
    </TabPane>
    <TabPane tabId="Monitor" class="mt-4">
        <span slot="tab" class="small">
            <Activity />
            Monitor
             <!-- <Icon name="alarm" /> -->
        </span>
        <Dashboard />
        
    </TabPane>
    <TabPane tabId="autostorage" class="mt-4">
        <span slot="tab" class="small">
            <Workflow />
            Auto Storage
             <!-- <Icon name="alarm" /> -->
        </span>
        <Dashboard />
        
    </TabPane>
    <TabPane tabId="settings" class="mt-4">
        <span slot="tab" class="small">
            <Wrench />
            Settings
             <!-- <Icon name="alarm" /> -->
        </span>
        <Dashboard />
        
    </TabPane>
</TabContent>
