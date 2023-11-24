export namespace main {
	
	export class Person {
	    name: string;
	    nickName: string;
	
	    static createFrom(source: any = {}) {
	        return new Person(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.nickName = source["nickName"];
	    }
	}

}

export namespace models {
	
	export class BaseRequest {
	    key: string;
	    value: string;
	    keyType: number;
	    action: number;
	    doByOpHistory: boolean;
	
	    static createFrom(source: any = {}) {
	        return new BaseRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	        this.keyType = source["keyType"];
	        this.action = source["action"];
	        this.doByOpHistory = source["doByOpHistory"];
	    }
	}
	export class BaseResponse {
	    code: number;
	    message: string;
	    data: any;
	    total: number;
	
	    static createFrom(source: any = {}) {
	        return new BaseResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	        this.total = source["total"];
	    }
	}
	export class GlobalConfigReq {
	    jsonFormat: boolean;
	    etcdEndPoint: string;
	
	    static createFrom(source: any = {}) {
	        return new GlobalConfigReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.jsonFormat = source["jsonFormat"];
	        this.etcdEndPoint = source["etcdEndPoint"];
	    }
	}
	export class ListRequest {
	    limit: number;
	
	    static createFrom(source: any = {}) {
	        return new ListRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.limit = source["limit"];
	    }
	}
	export class ReqJsonDiff {
	    old: string;
	    new: string;
	
	    static createFrom(source: any = {}) {
	        return new ReqJsonDiff(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.old = source["old"];
	        this.new = source["new"];
	    }
	}

}

