// adapter模式-适配器模式
// 适配器我们都用过，比方你出去旅游，国内的充电器插头和国外的不一样
// 这样就逼得你必须买一个适配器，适配器的一头和国外匹配，一头和国内的
// 匹配。
// 适配器模式有目标角色，源角色，适配器角色
// 目标角色就是我们希望得到的接口，源角色就是现在我们有的接口。适配器角色
// 就是需要完成将源接口转化为目标接口

<?php
/** 
 * 目标角色 
 * @version 2.0 
 */  
interface Target {  
    public function hello();  
    public function world();  
}  
   
/** 
 * 源角色：被适配的角色 
 */  
class Adaptee {  
    public function world() {  
        echo ' world <br />';  
    } 
    public function greet() {  
        echo ' Greet ';  
    }
}  
   
/** 
 * 类适配器角色 
 */  
class Adapter extends Adaptee implements Target {  
    public function hello() {  
       parent::greet();
    }  
   
}
/** 
 * 客户端程序 
 * 
 */  
class Client {  
   
    /** 
     * Main program. 
     */  
    public static function main() {  
        $adapter = new Adapter();  
        $adapter->hello();  
        $adapter->world();  
    }  
}  
Client::main();