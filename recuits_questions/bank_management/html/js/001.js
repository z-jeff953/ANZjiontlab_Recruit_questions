function custom_close(){
    if(confirm("您确定要关闭本页吗？")){
        window.opener=null;
        window.open('','_self');
        window.close();
    }
}
//关闭按钮
const btn = document.querySelector('#subbtn');

const nameInput = document.querySelector('#name');
nameInput.addEventListener('input', e => {
    console.log(nameInput.value);
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

const xInput = document.querySelector('#x');

xInput.addEventListener('input', e => {
    console.log(xInput.value);
});

const yInput = document.querySelector('#y');

yInput.addEventListener('input', e => {
    console.log(yInput.value);
});

btn.addEventListener('click',(e) => {
    e.preventDefault();
    console.log(e.target.className);
    document.getElementById('subbtn').style.background = '#ccc';
    var json1 ={"id":`${idInput.value}`,"name":`${nameInput.value}`,"password":`${pwInput.value}`,"money":`${moneyInput.value}`,"x":`${xInput.value}`,"y":`${yInput.value}`};
    js2 = JSON.stringify(json1);
    console.log(json1);
    console.log(js2);
    //学了一天没学会怎么把json文件返回后端处理,可以在控制台验证有数据输入
});