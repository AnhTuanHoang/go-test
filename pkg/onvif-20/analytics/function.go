
package analytics

type CreateAnalyticsModulesFunction struct{}

func (_ *CreateAnalyticsModulesFunction) Request() interface{} {
	return &CreateAnalyticsModules{}
}
func (_ *CreateAnalyticsModulesFunction) Response() interface{} {
	return &CreateAnalyticsModulesResponse{}
}

type CreateRulesFunction struct{}

func (_ *CreateRulesFunction) Request() interface{} {
	return &CreateRules{}
}
func (_ *CreateRulesFunction) Response() interface{} {
	return &CreateRulesResponse{}
}

type DeleteAnalyticsModulesFunction struct{}

func (_ *DeleteAnalyticsModulesFunction) Request() interface{} {
	return &DeleteAnalyticsModules{}
}
func (_ *DeleteAnalyticsModulesFunction) Response() interface{} {
	return &DeleteAnalyticsModulesResponse{}
}

type DeleteRulesFunction struct{}

func (_ *DeleteRulesFunction) Request() interface{} {
	return &DeleteRules{}
}
func (_ *DeleteRulesFunction) Response() interface{} {
	return &DeleteRulesResponse{}
}

type GetAnalyticsModuleOptionsFunction struct{}

func (_ *GetAnalyticsModuleOptionsFunction) Request() interface{} {
	return &GetAnalyticsModuleOptions{}
}
func (_ *GetAnalyticsModuleOptionsFunction) Response() interface{} {
	return &GetAnalyticsModuleOptionsResponse{}
}

type GetAnalyticsModulesFunction struct{}

func (_ *GetAnalyticsModulesFunction) Request() interface{} {
	return &GetAnalyticsModules{}
}
func (_ *GetAnalyticsModulesFunction) Response() interface{} {
	return &GetAnalyticsModulesResponse{}
}

type GetRuleOptionsFunction struct{}

func (_ *GetRuleOptionsFunction) Request() interface{} {
	return &GetRuleOptions{}
}
func (_ *GetRuleOptionsFunction) Response() interface{} {
	return &GetRuleOptionsResponse{}
}

type GetRulesFunction struct{}

func (_ *GetRulesFunction) Request() interface{} {
	return &GetRules{}
}
func (_ *GetRulesFunction) Response() interface{} {
	return &GetRulesResponse{}
}

type GetSupportedAnalyticsModulesFunction struct{}

func (_ *GetSupportedAnalyticsModulesFunction) Request() interface{} {
	return &GetSupportedAnalyticsModules{}
}
func (_ *GetSupportedAnalyticsModulesFunction) Response() interface{} {
	return &GetSupportedAnalyticsModulesResponse{}
}

type GetSupportedRulesFunction struct{}

func (_ *GetSupportedRulesFunction) Request() interface{} {
	return &GetSupportedRules{}
}
func (_ *GetSupportedRulesFunction) Response() interface{} {
	return &GetSupportedRulesResponse{}
}

type ModifyAnalyticsModulesFunction struct{}

func (_ *ModifyAnalyticsModulesFunction) Request() interface{} {
	return &ModifyAnalyticsModules{}
}
func (_ *ModifyAnalyticsModulesFunction) Response() interface{} {
	return &ModifyAnalyticsModulesResponse{}
}

type ModifyRulesFunction struct{}

func (_ *ModifyRulesFunction) Request() interface{} {
	return &ModifyRules{}
}
func (_ *ModifyRulesFunction) Response() interface{} {
	return &ModifyRulesResponse{}
}
