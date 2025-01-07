// example.h
#ifndef EXAMPLE_H
#define EXAMPLE_H

#ifdef __cplusplus
extern "C" {
#endif

    typedef struct {
        int intValue;
        double doubleValue;
    } Data;

    __declspec(dllexport) double processData(Data data);

#ifdef __cplusplus
}
#endif

#endif // EXAMPLE_H
