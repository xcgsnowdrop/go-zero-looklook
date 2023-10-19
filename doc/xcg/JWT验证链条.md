# JWT定义

JWT（JSON Web Token）是一种用于在网络中传递信息的开放标准（RFC 7519），它通常用于身份验证和信息交换。JWT是一个紧凑且自包含的数据结构，由三部分组成：头部（Header）、载荷（Payload）、和签名（Signature）。下面是JWT的用法详解：

### JWT结构：

一个JWT通常由三个部分组成，它们以Base64 URL编码后的字符串用点号 `.` 分隔：

1. **头部（Header）：** 头部通常由两部分组成，标识令牌类型和所使用的签名算法，如下所示：

   ```json
   {
       "alg": "HS256",
       "typ": "JWT"
   }
   ```

2. **载荷（Payload）：** 载荷包含有关令牌的有关信息，以及所需要传递的其他数据。有三种类型的声明：注册声明、公共声明和私有声明。载荷可以包含标准声明（例如，用户ID、过期时间等）和自定义声明。

   ```json
   {
       "sub": "1234567890",
       "name": "John Doe",
       "iat": 1516239022
   }
   ```

3. **签名（Signature）：** 签名用于验证令牌的完整性和验证令牌的来源。签名是使用头部中指定的签名算法和密钥对头部和载荷的组合进行计算得到的。

### 使用JWT进行认证和授权：

JWT可以用于认证用户和授权访问资源。以下是JWT的一般使用方式：

1. **用户登录：** 用户提供其凭证（通常是用户名和密码）进行登录。服务器验证凭证，如果验证通过，会生成一个JWT并将其返回给客户端。

2. **JWT的传递：** 客户端在后续请求中将JWT放在请求头、URL参数或Cookie中。

3. **验证和解析：** 服务器接收请求后，验证JWT的签名以确保其完整性，然后解析JWT以获取载荷中的信息，如用户ID、权限等。

4. **授权检查：** 服务器使用JWT中的信息进行访问控制检查，以确定用户是否有权执行请求的操作。

5. **响应：** 如果用户被授权，服务器处理请求并返回响应。

### JWT的优势：

- **自包含：** JWT 包含所有必要的信息，不需要在服务器端存储会话状态。
- **跨域支持：** JWT 可以在不同域之间传递信息，适用于单点登录 (SSO) 方案。
- **可扩展性：** 你可以添加自定义声明来传递额外信息。
- **安全性：** JWT 的签名可以保护令牌的完整性，确保其不被篡改。
- **快速解析：** JWT 可以很容易地被解码和解析，因此在验证令牌时效率高。

### 注意事项：

- 避免在JWT中存储敏感信息，因为JWT的内容可以被解码。敏感信息应该存储在服务器端。
- 使用足够强的签名算法和密钥，以防止令牌被篡改。
- 令牌的过期时间是重要的，应该合理设置以确保安全性。
- 谨慎处理跨站点请求伪造（CSRF）攻击，确保JWT不会被滥用。

# web项目中JWT验证流程

在基于浏览器的Web项目中，通常不会将JWT签名密钥直接暴露给Web用户。相反，JWT签名密钥通常保留在服务器端，而Web用户不需要知道密钥的具体值。以下是一种常见的工作流程：

1. **用户登录：** 当Web用户在浏览器中进行登录时，他们会提供用户名和密码等凭据。

2. **身份验证：** 服务器会验证用户提供的凭据，并如果验证成功，则生成JWT令牌。服务器会使用服务器端存储的密钥来签署JWT。

3. **JWT返回：** 服务器将生成的JWT令牌作为响应的一部分返回给Web用户的浏览器。Web用户将JWT令牌存储在浏览器的本地存储（例如，LocalStorage 或 Cookie）中。

4. **JWT的使用：** Web用户在随后的HTTP请求中将JWT令牌发送给服务器，以便进行授权。

5. **服务器验证：** 服务器接收到JWT令牌后，会使用服务器端存储的密钥来验证JWT的签名和完整性。

6. **授权：** 如果JWT有效并通过验证，服务器将使用JWT中的信息来授权用户对请求资源的访问。

在这个流程中，Web用户不需要知道服务器使用的JWT签名密钥。密钥是服务器的一部分，保密存储在服务器端，用于签署和验证JWT令牌。Web用户只需要在后续的请求中发送JWT令牌，服务器会使用自己的密钥来验证令牌。

这个流程的关键点是确保服务器的JWT签名密钥保持安全。如果密钥泄漏，JWT令牌可能会被篡改。因此，服务器的密钥管理是至关重要的，应采取适当的安全措施来保护密钥。

# go-zero中jwt验证流程
- app\usercenter\cmd\api\usercenter.go中 handler.RegisterHandlers(server, ctx)
- app\usercenter\cmd\api\internal\handler\routes.go中 server.AddRoutes()带上rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret)参数
- 带上上述参数后，server.ngin.routes.jwt中的{enable:true,secret:""}enable被设置为true，secret被设置为实际密钥
- 注册完路由以后，server.ngin.routes结构如下：
```json
(*server).ngin.routes: 
[
    github.com/zeromicro/go-zero/rest.featuredRoutes {
        timeout: 0,
        priority: false,
        jwt: github.com/zeromicro/go-zero/rest.jwtSetting {enabled: false, secret: "", prevSecret: ""},
        signature: github.com/zeromicro/go-zero/rest.signatureSetting {SignatureConf: {}, enabled: false}, 
        routes: [
            github.com/zeromicro/go-zero/rest.Route {
                Method: "POST", 
                Path: "/usercenter/v1/user/register", 
                Handler: looklook/app/usercenter/cmd/api/internal/handler/user.RegisterHandler.func1
            },
            github.com/zeromicro/go-zero/rest.Route {
                Method: "POST", 
                Path: "/usercenter/v1/user/login", 
                Handler: looklook/app/usercenter/cmd/api/internal/handler/user.LoginHandler.func1
            }  
        ]
    },
    github.com/zeromicro/go-zero/rest.featuredRoutes {
        timeout: 0, 
        priority: false, 
        jwt: github.com/zeromicro/go-zero/rest.jwtSetting {enabled: true, secret: "ae0536f9-6450-4606-8e13-5a19ed505da0", prevSecret: ""}, 
        signature: github.com/zeromicro/go-zero/rest.signatureSetting {SignatureConf: {}, enabled: false}, 
        routes: []github.com/zeromicro/go-zero/rest.Route len: 2, cap: 2, [
            {
                Method: "POST", 
                Path: "/usercenter/v1/user/detail", 
                Handler: looklook/app/usercenter/cmd/api/internal/handler/user.DetailHandler.func1
            },
            {
                Method: "POST", 
                Path: "/usercenter/v1/user/wxMiniAuth", 
                Handler: looklook/app/usercenter/cmd/api/internal/handler/user.WxMiniAuthHandler.func1
            }
        ]
    }
]
```
- 注意看上面把路由分成了两组rest.featuredRoutes，jwt.enabled=false的一组，jwt.enable=true的一组
- server.Start() ==> s.ngin.start() ==> ng.bindRoutes() ==> ng.bindFeaturedRoutes() ==> ng.bindRoute()该方法构建了一个中间件的调用链，并调用ng.appendAuthHandler()添加jwt校验处理器到中间件调用链中
``````go
func (ng *engine) appendAuthHandler(fr featuredRoutes, chn chain.Chain,
	verifier func(chain.Chain) chain.Chain) chain.Chain {
	if fr.jwt.enabled {
		if len(fr.jwt.prevSecret) == 0 {
			chn = chn.Append(handler.Authorize(fr.jwt.secret,
				handler.WithUnauthorizedCallback(ng.unauthorizedCallback)))
		} else {
			chn = chn.Append(handler.Authorize(fr.jwt.secret,
				handler.WithPrevSecret(fr.jwt.prevSecret),
				handler.WithUnauthorizedCallback(ng.unauthorizedCallback)))
		}
	}

	return verifier(chn)
}
``````
其中的handler.Authorize()指C:\Users\BFChainer\go\pkg\mod\github.com\zeromicro\go-zero@v1.5.3\rest\handler\authhandler.go中的
``````go
// Authorize returns an authorization middleware.
func Authorize(secret string, opts ...AuthorizeOption) func(http.Handler) http.Handler {
	var authOpts AuthorizeOptions
	for _, opt := range opts {
		opt(&authOpts)
	}

	parser := token.NewTokenParser()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tok, err := parser.ParseToken(r, secret, authOpts.PrevSecret)
			if err != nil {
				unauthorized(w, r, err, authOpts.Callback)
				return
			}

			if !tok.Valid {
				unauthorized(w, r, errInvalidToken, authOpts.Callback)
				return
			}

			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				unauthorized(w, r, errNoClaims, authOpts.Callback)
				return
			}

			ctx := r.Context()
			for k, v := range claims {
				switch k {
				case jwtAudience, jwtExpire, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
					// ignore the standard claims
				default:
					ctx = context.WithValue(ctx, k, v)
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
``````
 