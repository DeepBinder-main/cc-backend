package api


type Machine struct {
    Hostname  string `json:"hostname"`
    OsVersion string `json:"os_version"`
    IpAddress string `json:"ip_address"`
}
type MachineConf struct {
    MachineID  string `json:"machine_id"`
    Hostname   string `json:"hostname"`
    Username   string `json:"username"`
    Passphrase string `json:"passphrase,omitempty"`
    PortNumber int32  `json:"port_number"`
    Password   string `json:"password,omitempty"`
    HostKey    string `json:"host_key,omitempty"`
    FolderPath string `json:"folder_path,omitempty"`
}
type RabbitMqConfig struct {
    ConnUrl  string `json:"conn_url"`
    Username string `json:"username"`
    Password string `json:"password"`
}
type InfluxdbConfiguration struct {
    Type                 string `json:"type"`
    DatabaseName         string `json:"database_name"`
    Host                 string `json:"host"`
    Port                 int32  `json:"port"`
    User                 string `json:"user"`
    Password             string `json:"password"`
    Organization         string `json:"organization"`
    SslEnabled           bool   `json:"ssl_enabled"`
    BatchSize            int32  `json:"batch_size"`
    RetryInterval        string `json:"retry_interval"`
    RetryExponentialBase int32  `json:"retry_exponential_base"`
    MaxRetries           int32  `json:"max_retries"`
    MaxRetryTime         string `json:"max_retry_time"`
    MetaAsTags           string `json:"meta_as_tags,omitempty"`
}
type FileStashUrl struct {
    Url string `json:"url"`
}
type LvStorageIssuer struct {
    MachineID           string  `json:"machine_id"`
    IncBuffer           int     `json:"inc_buffer,omitempty"`
    DecBuffer           int     `json:"dec_buffer,omitempty"`
    Hostname            string  `json:"hostname"`
    Username            string  `json:"username"`
    MinAvailableSpaceGB float64 `json:"min_available_space_gb"`
    MaxAvailableSpaceGB float64 `json:"max_available_space_gb"`
}
type Notification struct {
    ID      string `json:"id"`
    Message string `json:"message"`
    CreatedAt string `json:"created_at"`
}
type RealtimeLog struct {
    LogMessage string `json:"log_message"`
    MachineID  string `json:"machine_id"`
    Timestamp  string `json:"timestamp"`
}
type VolumeGroup struct {
    MachineID string `json:"machine_id"`
    VgName    string `json:"vg_name"`
    PvCount   string `json:"pv_count"`
    LvCount   string `json:"lv_count"`
    SnapCount string `json:"snap_count"`
    VgAttr    string `json:"vg_attr"`
    VgSize    string `json:"vg_size"`
    VgFree    string `json:"vg_free"`
}
type PhysicalVolume struct {
    MachineID string `json:"machine_id"`
    PvName    string `json:"pv_name"`
    VgName    string `json:"vg_name"`
    PvFmt     string `json:"pv_fmt"`
    PvAttr    string `json:"pv_attr"`
    PvSize    string `json:"pv_size"`
    PvFree    string `json:"pv_free"`
}
type LogicalVolume struct {
    MachineID string `json:"machine_id"`
    LvName    string `json:"lv_name"`
    VgName    string `json:"vg_name"`
    LvAttr    string `json:"lv_attr"`
    LvSize    string `json:"lv_size"`
}