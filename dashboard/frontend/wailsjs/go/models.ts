export namespace platform {
	
	export class SystemStats {
	    id: number;
	    device_id: number;
	    mac_address: string;
	    hostname: string;
	    platform: string;
	    cpu: number;
	    ram: number;
	    disk: number;
	    incoming_net_traffic: number;
	    outbound_net_traffic: number;
	    processes: string;
	    time: number;
	
	    static createFrom(source: any = {}) {
	        return new SystemStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.device_id = source["device_id"];
	        this.mac_address = source["mac_address"];
	        this.hostname = source["hostname"];
	        this.platform = source["platform"];
	        this.cpu = source["cpu"];
	        this.ram = source["ram"];
	        this.disk = source["disk"];
	        this.incoming_net_traffic = source["incoming_net_traffic"];
	        this.outbound_net_traffic = source["outbound_net_traffic"];
	        this.processes = source["processes"];
	        this.time = source["time"];
	    }
	}
	export class Alert {
	    id: number;
	    device_id: number;
	    time: number;
	    hostname: string;
	    type: string;
	    value: number;
	    threshold: number;
	    severity: number;
	    system_stats_id: number;
	    system_stats: SystemStats;
	
	    static createFrom(source: any = {}) {
	        return new Alert(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.device_id = source["device_id"];
	        this.time = source["time"];
	        this.hostname = source["hostname"];
	        this.type = source["type"];
	        this.value = source["value"];
	        this.threshold = source["threshold"];
	        this.severity = source["severity"];
	        this.system_stats_id = source["system_stats_id"];
	        this.system_stats = this.convertValues(source["system_stats"], SystemStats);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Device {
	    id: number;
	    mac_address: string;
	    hostname: string;
	    platform: string;
	    system_stats: SystemStats[];
	    alerts: Alert[];
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.mac_address = source["mac_address"];
	        this.hostname = source["hostname"];
	        this.platform = source["platform"];
	        this.system_stats = this.convertValues(source["system_stats"], SystemStats);
	        this.alerts = this.convertValues(source["alerts"], Alert);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

