#include <std/math.pat>
#include <std/mem.pat>
#include <std/io.pat>

struct file_header {
  u32 magic;
  u32 unk04;
  u32 ts1;
  u32 ts2;
  s32 neg;
  char by_name1[8];
  padding[8];
  char by_name2[8];
  padding[8];
  char BUILD[8];
  padding[8];
};

struct inode {
  padding[24];
    
  u32 data_sz_1;
  
  u32 ts1;
  u32 ts2;
  u32 ts3;
  char id1[4];
  char id2[4];
  char id3[4];
  char id4[4];
  
  u32 data_sz_2;
  
  padding[4];
  char name[];
  padding[std::mem::align_to(8,$)-$];

  u32 unk1;
  u32 unk2;
  u32 unk3;
  u32 unk4;
  u32 unk5;
  u32 unk6;
  u32 unk7;
  u32 unk8;
  u32 unk9;
  u32 unk10;
  
  std::print("File {} at {:#x}..{:#x} sz {:#x}", name, $, $+data_sz_2, data_sz_2);

  u8  data[data_sz_2];
  padding[std::mem::align_to(8,$)-$];
};

file_header header @ 0x00;
//inode entries[while(!std::mem::eof())] @ 0x48;
inode entries[100] @ 0x48;
