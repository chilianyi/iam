# OpenPitrix中的IAM服务

iam模块本身尽量剥离了OpenPitrix相关的业务，因此iam模块本书是业务无关独立使用的。

OpenPitrix在iam模块的基础之上可以构建业务相关的账户和权限管理模块。
OpenPitrix的账号和权限管理模块需要在op仓库创建新的grpc服务接口，并针对业务增加胶水代码。

接口需要调整部分，组和用户管理，用户授权基于资源空间和方法（隐藏rbac概念）。有三类权限admin/isv_admin/user是写死的，直接在创建时指定。默认的角色和默认的资源空间信息在扩展信息字段制定。

增加一个根据组ID，获取组名字空间的方法，返回的是名字空间和组织路径(namespace和org_path)。
