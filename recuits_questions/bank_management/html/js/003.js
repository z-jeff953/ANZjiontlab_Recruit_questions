function custom_close(){
    if(confirm("您确定要关闭本页吗？")){
        window.opener=null;
        window.open('','_self');
        window.close();
    }
}
//关闭按钮
const btn = document.querySelector('#subbtn');

const targetInput = document.querySelector('#target');
targetInput.addEventListener('input', e => {
    console.log(targetInput.value);
});

const idInput = document.querySelector('#id');
idInput.addEventListener('input', e => {
    console.log(idInput.value);
});

const pwInput = document.querySelector('#password');
pwInput.addEventListener('input', e => {
    console.log(pwInput.value);
});

const moneyInput = document.querySelector('#money');

moneyInput.addEventListener('input', e => {
    console.log(moneyInput.value);
});

btn.addEventListener('click',(e) => {
    e.preventDefault();
    console.log(e.target.className);
    document.getElementById('subbtn').style.background = '#ccc';
    var json1 ={"id":`${idInput.value}`,"target":`${targetInput.value}`,"password":`${pwInput.value}`,"money":`${moneyInput.value}`};
    js2 = JSON.stringify(json1);
    console.log(json1);
    console.log(js2);
    //学了一天没学会怎么把json文件返回后端处理,可以在控制台验证有数据输入
});