// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type ForeverProcess struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*ForeverProcess
	namer   func(string) string
	connID  int
	
	Id         	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name       	string  	`db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Command    	string  	`db:"command" bson:"command" comment:"命令" json:"command" xml:"command"`
	Workdir    	string  	`db:"workdir" bson:"workdir" comment:"工作目录" json:"workdir" xml:"workdir"`
	Env        	string  	`db:"env" bson:"env" comment:"环境变量" json:"env" xml:"env"`
	Args       	string  	`db:"args" bson:"args" comment:"命令参数" json:"args" xml:"args"`
	Pidfile    	string  	`db:"pidfile" bson:"pidfile" comment:"PID记录文件" json:"pidfile" xml:"pidfile"`
	Logfile    	string  	`db:"logfile" bson:"logfile" comment:"日志记录文件" json:"logfile" xml:"logfile"`
	Errfile    	string  	`db:"errfile" bson:"errfile" comment:"错误记录文件" json:"errfile" xml:"errfile"`
	Respawn    	uint    	`db:"respawn" bson:"respawn" comment:"重试次数(进程被外部程序结束后自动启动)" json:"respawn" xml:"respawn"`
	Delay      	string  	`db:"delay" bson:"delay" comment:"延迟启动(例如1ms/1s/1m/1h)" json:"delay" xml:"delay"`
	Ping       	string  	`db:"ping" bson:"ping" comment:"心跳时间(例如1ms/1s/1m/1h)" json:"ping" xml:"ping"`
	Pid        	int     	`db:"pid" bson:"pid" comment:"PID" json:"pid" xml:"pid"`
	Status     	string  	`db:"status" bson:"status" comment:"进程运行状态" json:"status" xml:"status"`
	Debug      	string  	`db:"debug" bson:"debug" comment:"DEBUG" json:"debug" xml:"debug"`
	Disabled   	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated    	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Error      	string  	`db:"error" bson:"error" comment:"错误信息" json:"error" xml:"error"`
	Lastrun    	uint    	`db:"lastrun" bson:"lastrun" comment:"上次运行时间" json:"lastrun" xml:"lastrun"`
	Description	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
}

func (this *ForeverProcess) Trans() *factory.Transaction {
	return this.trans
}

func (this *ForeverProcess) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *ForeverProcess) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *ForeverProcess) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *ForeverProcess) Objects() []*ForeverProcess {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *ForeverProcess) NewObjects() *[]*ForeverProcess {
	this.objects = []*ForeverProcess{}
	return &this.objects
}

func (this *ForeverProcess) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *ForeverProcess) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *ForeverProcess) Name_() string {
	if this.namer != nil {
		return this.namer("forever_process")
	}
	return factory.TableNamerGet("forever_process")(this)
}

func (this *ForeverProcess) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *ForeverProcess) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *ForeverProcess) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *ForeverProcess) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *ForeverProcess) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *ForeverProcess) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *ForeverProcess) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *ForeverProcess) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Debug) == 0 { this.Debug = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Status) == 0 { this.Status = "idle" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *ForeverProcess) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *ForeverProcess) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *ForeverProcess) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["debug"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["debug"] = "N" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	if val, ok := kvset["status"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["status"] = "idle" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *ForeverProcess) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Debug) == 0 { this.Debug = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Status) == 0 { this.Status = "idle" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *ForeverProcess) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *ForeverProcess) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *ForeverProcess) Reset() *ForeverProcess {
	this.Id = 0
	this.Name = ``
	this.Command = ``
	this.Workdir = ``
	this.Env = ``
	this.Args = ``
	this.Pidfile = ``
	this.Logfile = ``
	this.Errfile = ``
	this.Respawn = 0
	this.Delay = ``
	this.Ping = ``
	this.Pid = 0
	this.Status = ``
	this.Debug = ``
	this.Disabled = ``
	this.Created = 0
	this.Updated = 0
	this.Error = ``
	this.Lastrun = 0
	this.Description = ``
	return this
}

func (this *ForeverProcess) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["Command"] = this.Command
	r["Workdir"] = this.Workdir
	r["Env"] = this.Env
	r["Args"] = this.Args
	r["Pidfile"] = this.Pidfile
	r["Logfile"] = this.Logfile
	r["Errfile"] = this.Errfile
	r["Respawn"] = this.Respawn
	r["Delay"] = this.Delay
	r["Ping"] = this.Ping
	r["Pid"] = this.Pid
	r["Status"] = this.Status
	r["Debug"] = this.Debug
	r["Disabled"] = this.Disabled
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["Error"] = this.Error
	r["Lastrun"] = this.Lastrun
	r["Description"] = this.Description
	return r
}

