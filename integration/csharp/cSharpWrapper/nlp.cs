using System;
using System.Runtime.InteropServices;
using System.Text;

public class MainClass
{
    public struct GoSlice
    {
        public IntPtr data;
        public long len,
            cap;

        public GoSlice(IntPtr data, long len, long cap)
        {
            this.data = data;
            this.len = len;
            this.cap = cap;
        }
    }

    public struct GoString
    {
        public string msg;
        public long len;

        public GoString(string msg, long len)
        {
            this.msg = msg;
            this.len = len;
        }
    }

    [DllImport(
        "libgo.dll",
        EntryPoint = "GetLabels",
        CharSet = CharSet.Unicode,
        CallingConvention = CallingConvention.StdCall
    )]
    public static extern string GetLabels(GoString inputText); // function signature, must have a return type

    [DllImport("libgo.dll", EntryPoint = "TokeniseThis")]
    public static extern void TokeniseThis();

    [DllImport("libgo.dll", EntryPoint = "Sum")]
    public static extern int Sum(int a, int b);

    [DllImport("libgo.dll", EntryPoint = "stringtest")]
    extern static int stringtest(byte[] test);

    static void Main()
    {
        string msg = "Tarun is a 25 yrs old boy and he has $10000000 in his bank";

        GoString str = new GoString(msg, msg.Length);
        // /*  string[] arr =  */GetLabels(
        // Encoding.ASCII.GetBytes("Tarun is a 25 yrs old boy and he has $10000000 in his bank")
        // );
        // string[] arr = { "arr", "ee" };
        // AsSpan(arr);
        // Console.WriteLine(arr[0]);
        GetLabels(str);
        TokeniseThis();
        Console.WriteLine(Sum(1, 2));
        stringtest(Encoding.ASCII.GetBytes("I Am String"));
    }

    public static void AsSpan(string[] array)
    {
        foreach (var item in array)
        {
            Console.Write($" ==> {item}");
        }
    }
}
