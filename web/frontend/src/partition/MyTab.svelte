<script>
    import { TabContent, TabPane } from "sveltestrap";
    import DashboardTab from "./tabs/Dashboard.svelte"
    import CreateLVMTab from "./tabs/CreateLVM.svelte"
    import MonitorTab from "./tabs/Monitor.svelte"
    import AutoStorageTab from "./tabs/AutoStorage.svelte"
    import SettingTab from "./tabs/Setting.svelte"
    import AutoStorage from "./tabs/AutoStorage.svelte";
</script>

<TabContent vertical pills class="p-3">
    <TabPane tabId="dashboard" active>
        <span class="fs-4" slot="tab">
            Dashboard
        </span>
        <DashboardTab />
    </TabPane>
    <TabPane tabId="create-lvm">
        <span class="fs-4" slot="tab">
            Create LVM
        </span>

        <CreateLVMTab />
    </TabPane>
    
    <TabPane tabId="monitor-tab">
        <span class="fs-4" slot="tab">
            Monitor
        </span>

        <MonitorTab />
    </TabPane>
    
    <TabPane tabId="auto-storage">
        <span class="fs-4" slot="tab">
            Auto Storage
        </span>

        <AutoStorage />
    </TabPane>
    
    <TabPane tabId="setting-tab">
        <span class="fs-4" slot="tab">
            Setting
        </span>

        <SettingTab />
    </TabPane>
    
</TabContent>
