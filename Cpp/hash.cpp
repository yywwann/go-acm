#include <iostream>
#include <string>
#include <algorithm>
#include <vector>

using std::cin; using std::cout;
using std::endl; using std::vector;
using std::reverse; using std::max;
using std::string; using std::min;

typedef long long LL;
const int P = 10000019;//计算hash值的进制数
const int MOD = 1000000007;//计算hash值的模数
const int MAXN = 200010;//字符串最大长度
LL PowP[MAXN], H1[MAXN], H2[MAXN];//PowP[i]存放p^i%MOD,H1、H2分别存放str1和str2的hash

void Init(int len){//初始化PowP
    PowP[0] = 1;
    for (int i = 1; i <= len; ++i)
    {
        PowP[i] = (PowP[i -1] * P) % MOD;
    }
}

//计算字符串str的hash
void calH(LL H[], string &str){
    H[0] = str[0];
    for (int i = 1; i < str.length(); ++i)
    {
        H[i] = (H[i - 1] * P + str[i]) % MOD;
    }
}

//计算H[i...j]
int calSingleSubH(LL H[], int i , int j){
    if(i == 0) return H[j];
    else return ((H[j] - H[i-1] * PowP[j - i + 1])% MOD + MOD) % MOD;
}

//在[l,r]里二分回文半径；len：字符串长；i：对称点；
//isEven：求奇回文时为0，偶回文为1；
//寻找最后一个满足"hashL==hashR"的回文半径
//等价于寻找第一个满足条件的"hashL != hashR"的回文半径，减1
int binartSearch(int l, int r, int len, int i, int isEven){
    while(l < r){
        int mid = (r + l) / 2;
        int H1L = i - mid + isEven, H1R = i;
        int H2L = len - 1 - (i + mid), H2R = len - 1 - (i + isEven);
        int hashL = calSingleSubH(H1, H1L, H1R);
        int hashR = calSingleSubH(H2, H2L, H2R);

        if(hashL != hashR){
            r = mid;
        }else{
            l = mid + 1;
        }
    }

    return l - 1;
}


int main(int argc, char const *argv[])
{
    string str;
    getline(cin, str);

    Init(str.length());

    calH(H1, str);

    reverse(str.begin(), str.end());
    calH(H2, str);

    int ans = 0;
    //奇回文
    for (int i = 0; i < str.length(); ++i)
    {
        //二分上界为分界点i的左右长度较小值加1
        //回文半径的右边界，防止回文半径长度超过字符串长度
        int maxLen = min(i, (int)str.length() - 1 - i) + 1;
        int k = binartSearch(0, maxLen, str.length(), i, 0);
        ans = max(ans, k * 2 + 1);
    }
    //偶回文
    for (int i = 0; i < str.length(); ++i)
    {
        //二分上界为分界点i的左右长度较小值加1（注意：左长为i+1）
        int maxLen = min(i + 1, (int)str.length() - 1 - i) + 1;
        int k = binartSearch(0, maxLen, str.length(), i, 1);
        ans = max(ans, k * 2);
    }

    printf("%d\n", ans);
    return 0;
}
