function custom_close(){
    if(confirm("您确定要关闭本页吗？")){
        window.opener=null;
        window.open('','_self');
        window.close();
    }
}
//与001.js同理
const btn = document.querySelector('#subbtn');

const idInput = document.querySelector('#id');
idInput.addEventListener('input', e => {
    console.log(idInput.value);
});

const pwInput = document.querySelector('#password');
pwInput.addEventListener('input', e => {
    console.log(pwInput.value);
});

btn.addEventListener('click',(e) => {
    e.preventDefault();
    console.log(e.target.className);
    document.getElementById('subbtn').style.background = '#ccc';
    var json1 ={"id":`${idInput.value}`,"password":`${pwInput.value}`};
    js2 = JSON.stringify(json1);
    console.log(json1);
    console.log(js2);
});