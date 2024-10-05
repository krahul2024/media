package routes

func Register(router *Router){
    userRoutes(router.Group("/user"));
    postRoutes(router.Group("/post"));
}
